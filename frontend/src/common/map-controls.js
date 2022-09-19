/*

Copyright (c) 2022 Krassimir Valev

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

    PhotoPrism® is a registered trademark of Michael Mayer.  You may use it as required
    to describe our software, run your own server, for educational purposes, but not for
    offering commercial goods, products, or services without prior written permission.
    In other words, please ask.

Feel free to send an e-mail to hello@photoprism.org if you have questions,
want to support our work, or just want to say hello.

Additional information can be found in our Developer Guide:
https://docs.photoprism.org/developer-guide/

*/
import { DateTime } from "luxon";
import maplibregl from "maplibre-gl";

import Notify from "common/notify";
import { $gettext } from "common/vm";
import { Photo } from "model/photo";
import { CountriesTimeZones, PreferredTimeZones } from "options/options";

const SOURCE = "locate-nearby";

export class LocateNearbyControl {
  constructor({ model, settings, range, closest = 10 }) {
    this._model = model;
    this._settings = settings;
    this._range = range;
    this._closestN = closest;
  }

  onAdd(map) {
    this._map = map;
    this._container = this.createContainer();

    this._map.on("load", () => this.addLayers());

    return this._container;
  }

  onRemove() {
    this._container.parentNode.removeChild(this._container);
    this._map = undefined;
  }

  createContainer() {
    var radarIcon = document.createElement("span");
    radarIcon.className = "maplibregl-ctrl-icon mapboxgl-ctrl-icon material-icons md-dark";
    radarIcon.textContent = "radar";

    var radarButton = document.createElement("button");
    radarButton.className = "maplibregl-ctrl-icon mapboxgl-ctrl-icon";
    radarButton.style = "padding-top: 3px";
    radarButton.title = "Load timewise nearby photos";
    radarButton.addEventListener("click", (e) => this.onRadarClick(e));
    radarButton.appendChild(radarIcon);

    var clearIcon = document.createElement("span");
    clearIcon.className = "maplibregl-ctrl-icon mapboxgl-ctrl-icon material-icons md-dark";
    clearIcon.textContent = "delete";

    var clearButton = document.createElement("button");
    clearButton.className = "maplibregl-ctrl-icon mapboxgl-ctrl-icon";
    clearButton.style = "padding-top: 3px";
    clearButton.title = "Clear shown nearby photos";
    clearButton.addEventListener("click", (e) => this.onClearClick(e));
    clearButton.appendChild(clearIcon);

    const container = document.createElement("div");
    container.className = "maplibregl-ctrl maplibregl-ctrl-group mapboxgl-ctrl mapboxgl-ctrl-group";
    container.addEventListener("contextmenu", (e) => e.preventDefault());
    container.appendChild(radarButton);
    container.appendChild(clearButton);

    return container;
  }

  addLayers() {
    this._map.addSource(SOURCE, {
      data: {
        type: "FeatureCollection",
        features: [],
      },
      type: "geojson",
    });

    this._map.addLayer({
      id: SOURCE,
      source: SOURCE,
      type: "circle",
      filter: ["all", ["==", "$type", "Point"]],
      paint: {
        "circle-color": "hsla(200,80%,70%,0.5)",
        "circle-radius": 6,
        "circle-stroke-width": 2,
        "circle-stroke-color": "hsl(200,80%,50%)",
      },
    });

    this._map.on("click", SOURCE, (e) => {
      this._map.fire("radar.selected", {
        location: e.lngLat,
      });
    });

    // Create a popup, but dont add it to the map yet.
    const popup = new maplibregl.Popup({
      closeButton: false,
      closeOnClick: false,
      className: "plain-marker",
    });

    this._map.on("mouseover", SOURCE, (e) => {
      this._map.getCanvas().style.cursor = "pointer";

      // Copy coordinates array.
      const coordinates = e.features[0].geometry.coordinates.slice();
      const thumbnail = e.features[0].properties.thumbnail;

      // Ensure that if the map is zoomed out such that multiple
      // copies of the feature are visible, the popup appears
      // over the copy being pointed to.
      while (Math.abs(e.lngLat.lng - coordinates[0]) > 180) {
        coordinates[0] += e.lngLat.lng > coordinates[0] ? 360 : -360;
      }

      // Create the popup element showing the thumbnail
      let el = document.createElement("div");
      el.className = "marker";
      el.style.backgroundImage = `url(${thumbnail})`;
      el.style.width = "100px";
      el.style.height = "100px";

      // Populate the popup and set its coordinates based on the feature found.
      popup.setLngLat(coordinates).setDOMContent(el).addTo(this._map);
    });

    this._map.on("mouseleave", SOURCE, () => {
      this._map.getCanvas().style.cursor = "";
      popup.remove();
    });
  }

  removeLayers() {
    if (this._map.getSource(SOURCE)) {
      this._map.removeSource(SOURCE);
    }
  }

  onRadarClick(e) {
    e.preventDefault();

    const photoDate = this.guesstimateUtcDate(new Photo(this._model));
    const before = photoDate.plus(this._range);
    const after = photoDate.minus(this._range);

    const params = {
      count: Photo.limit(),
      offset: 0,
      merged: true,
      geo: true,
      beforet: before.toFormat("yyyy-LL-dd HH:mm:ss"),
      aftert: after.toFormat("yyyy-LL-dd HH:mm:ss"),
    };

    Photo.search(params).then((response) => {
      const features = response.data
        .filter((media) => media.Lng !== 0 && media.Lat !== 0)
        .sort(
          (a, b) =>
            Math.abs(photoDate - new Photo(a).utcDate()) -
            Math.abs(photoDate - new Photo(b).utcDate())
        )
        .slice(0, this._closestN)
        .map((media) => ({
          type: "Feature",
          geometry: {
            type: "Point",
            coordinates: [media.Lng, media.Lat],
          },
          properties: {
            year: media.Year,
            thumbnail: new Photo(media).thumbnailUrl("tile_100"),
          },
        }));

      this._map.getSource(SOURCE).setData({
        type: "FeatureCollection",
        features: features,
      });

      const coordinates = features.map((f) => f.geometry.coordinates);

      if (coordinates.length > 0) {
        // Create a "LngLatBounds" with both corners at the first coordinate.
        const bounds = new maplibregl.LngLatBounds(coordinates[0], coordinates[0]);

        // Extend the "LngLatBounds" to include every coordinate in the bounds result.
        for (const coord of coordinates) {
          bounds.extend(coord);
        }

        this.zoomTo(bounds, 17);
      }

      if (coordinates.length == 0) {
        Notify.info($gettext("No nearby photos"));
      }
    });
  }

  onClearClick(e) {
    e.preventDefault();

    this._map.getSource(SOURCE).setData({
      type: "FeatureCollection",
      features: [],
    });

    this.zoomTo(new maplibregl.LngLatBounds([0, 0], [0, 0]), 1);
  }

  zoomTo(bounds, maxZoom) {
    this._map.fitBounds(bounds, {
      maxZoom: maxZoom,
      padding: 100,
      duration: this._settings.animate,
      essential: false,
      animate: this._settings.animate > 0,
    });
  }

  guesstimateUtcDate(photo) {
    if (!photo.hasTimeZone() && photo.hasCountry()) {
      const country = photo.Country.toUpperCase();
      const ct = CountriesTimeZones();

      if (country in ct) {
        const timezones = ct[country];

        let timezone;

        if (timezones.length == 1) {
          timezone = timezones[0];
        } else if (timezones.length > 1) {
          // Some countries have more than one timezone, so try to figure out the timezone based on
          // the photo keywords, which might contain the city name.
          // The list of countries with more than one timezone can be replicated with this snippet:
          // https://runkit.com/embed/2yy43dnz45u3
          const keywords = photo.Details.Keywords.split(",")
            .map((x) => x.trim())
            .map((x) => x[0].toUpperCase() + x.slice(1));

          timezone = timezones.find((tz) => keywords.some((kw) => tz.mainCities.includes(kw)));

          // As a last resort fallback to the hardcoded preferred timezones.
          if (!timezone) {
            timezone = PreferredTimeZones()[country];
          }
        }

        if (timezone) {
          let iso = photo.localDateString();
          const zone = timezone.name || timezone;

          return DateTime.fromISO(iso, { zone }).toUTC();
        }
      }
    }

    // Either there is a timezone or there is nothing we can do.
    return photo.utcDate();
  }
}
