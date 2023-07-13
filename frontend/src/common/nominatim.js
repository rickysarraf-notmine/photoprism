/*

Copyright (c) 2021 Krassimir Valev

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

import Axios from "axios";

const api = Axios.create({
  baseURL: "https://nominatim.openstreetmap.org/",
  params: {
    format: "geojson",
  },
});

api.interceptors.response.use(function (response) {
  const data = response.data;

  // add required Carmen GeoJson properties
  data.attribution = data.license;
  data.query = [response.config.params.q];

  data.features.forEach((feature) => {
    feature.id = feature.properties.category + "." + feature.properties.osm_id;
    feature.place_name = feature.properties.display_name;
    feature.place_type = [feature.properties.category];
    feature.center = feature.geometry.coordinates;
  });

  return data;
});

const nominatim = {
  forwardGeocode: async (config) => {
    const params = {
      q: config.query,
      limit: config.limit,
    };

    return await api.get("/search", { params });
  },
  reverseGeocode: async (config) => {
    const params = {
      lat: config.query[0],
      lon: config.query[1],
    };

    return await api.get("/reverse", { params });
  },
};

export default nominatim;
