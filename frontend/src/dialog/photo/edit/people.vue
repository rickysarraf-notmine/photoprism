<template>
  <div class="p-tab p-tab-photo-people">
    <v-container grid-list-xs fluid class="pa-2 p-faces">
      <v-alert
          :value="markers.length === 0"
          color="secondary-dark" icon="lightbulb_outline" class="no-results ma-2 opacity-70" outline
      >
        <h3 class="body-2 ma-0 pa-0">
          <translate>No people found</translate>
        </h3>
        <p class="body-1 mt-2 mb-0 pa-0">
          <translate>You may rescan your library to find additional faces.</translate>
          <translate>Recognition starts after indexing has been completed.</translate>
        </p>
      </v-alert>
      <v-layout row wrap class="search-results face-results cards-view">
        <v-flex
            v-for="(marker, index) in markers"
            :key="index"
            xs12 sm6 md3 xl2 d-flex
        >
          <v-card tile
                  :data-id="marker.UID"
                  style="user-select: none;"
                  :class="marker.classes()"
                  class="result card">
            <div class="card-background card"></div>
            <v-img :src="marker.thumbnailUrl('tile_320')"
                   :transition="false"
                   aspect-ratio="1"
                   class="card darken-1">
              <v-btn v-if="!marker.UID" :ripple="false" :depressed="false" class="input-confirm-face"
                     icon flat small absolute :title="$gettext('Confirm face')"
                     @click.stop.prevent="onConfirm(marker, index)">
                <v-icon color="white" class="action-confirm-face">library_add_check</v-icon>
              </v-btn>
              <v-btn v-if="marker.UID && !marker.SubjUID && !marker.Invalid" :ripple="false" :depressed="false" class="input-reject"
                     icon flat small absolute :title="$gettext('Remove')"
                     @click.stop.prevent="onReject(marker)">
                <v-icon color="white" class="action-reject">clear</v-icon>
              </v-btn>
            </v-img>

            <v-card-actions class="card-details pa-0">
              <v-layout v-if="marker.Invalid" row wrap align-center>
                <v-flex xs12 class="text-xs-center pa-0">
                  <v-btn color="transparent" :disabled="busy"
                         large depressed block :round="false"
                         class="action-undo text-xs-center"
                         :title="$gettext('Undo')" @click.stop="onApprove(marker)">
                    <v-icon dark>undo</v-icon>
                  </v-btn>
                </v-flex>
              </v-layout>
              <v-layout v-else-if="!marker.UID" row wrap align-center>
                <v-flex xs12 class="text-xs-center pa-0">
                  <v-btn color="transparent" disabled
                         large depressed block
                         class="text-xs-center">
                  </v-btn>
                </v-flex>
              </v-layout>
              <v-layout v-else-if="marker.SubjUID" row wrap align-center>
                <v-flex xs12 class="text-xs-left pa-0">
                  <v-text-field
                      v-model="marker.Name"
                      :rules="[textRule]"
                      :disabled="busy"
                      :readonly="true"
                      browser-autocomplete="off"
                      autocorrect="off"
                      class="input-name pa-0 ma-0"
                      hide-details
                      single-line
                      solo-inverted
                      clearable
                      clear-icon="eject"
                      @click:clear="onClearSubject(marker)"
                      @change="onRename(marker)"
                      @keyup.enter.native="onRename(marker)"
                  ></v-text-field>
                </v-flex>
              </v-layout>
              <v-layout v-else row wrap align-center>
                <v-flex xs12 class="text-xs-left pa-0">
                  <v-combobox
                      v-model="marker.Name"
                      style="z-index: 250"
                      :items="$config.values.people"
                      item-value="Name"
                      item-text="Name"
                      :disabled="busy"
                      :return-object="false"
                      :menu-props="menuProps"
                      :allow-overflow="false"
                      :hint="$gettext('Name')"
                      hide-details
                      single-line
                      solo-inverted
                      open-on-clear
                      hide-no-data
                      append-icon=""
                      prepend-inner-icon="person_add"
                      browser-autocomplete="off"
                      class="input-name pa-0 ma-0"
                      @change="onRename(marker)"
                      @keyup.enter.native="onRename(marker)"
                  >
                  </v-combobox>
                </v-flex>
              </v-layout>
            </v-card-actions>
          </v-card>
        </v-flex>
      </v-layout>
      <div class="text-xs-center mt-3 mb-2">
        <v-btn
            color="secondary" round
            @click.stop="loadAllFaces()"
        >
          <translate>Load low quality faces</translate>
        </v-btn>
        <v-btn
            color="secondary" round
            @click.stop="dialog.select = true"
        >
          <translate>Manually select face</translate>
        </v-btn>
      </div>
    </v-container>
    <p-people-select-face-dialog
        lazy
        :show="dialog.select"
        :file="mainFile"
        @cancel="dialog.select = false"
        @confirm="onSelect"
    >
    </p-people-select-face-dialog>
  </div>
</template>

<script>
import Notify from "common/notify";
import Util from "common/util";
import File from "model/file";
import Marker from "model/marker";

const MarkerOverlapThreshold = 0.5;

export default {
  name: 'PTabPhotoPeople',
  props: {
    model: {
      type: Object,
      default: () => {},
    },
    uid: String,
  },
  data() {
    return {
      busy: false,
      markers: this.model.getMarkers(true),
      imageUrl: this.model.thumbnailUrl("fit_720"),
      mainFile: new File(this.model.mainFile()),
      disabled: !this.$config.feature("edit"),
      config: this.$config.values,
      readonly: this.$config.get("readonly"),
      menuProps:{"closeOnClick":false, "closeOnContentClick":true, "openOnClick":false, "maxHeight":300},
      textRule: (v) => {
        if (!v || !v.length) {
          return this.$gettext("Name");
        }

        return v.length <= this.$config.get('clip') || this.$gettext("Name too long");
      },
      dialog: {
        select: false,
      },
    };
  },
  methods: {
    refresh() {
    },
    markerFromFace(face) {
      // Create a Marker object from the backend "Face" object.
      // NB: For some reason y is used for x calculation and vice versa.
      const x = (face.face.y - face.face.size / 2) / face.cols;
      const y = (face.face.x - face.face.size / 2) / face.rows;
      const w = face.face.size / face.cols;
      const h = face.face.size / face.rows;

      const hex = (v) => parseInt(v * 1000).toString(16).padStart(3, '0');
      const area = `${hex(x)}${hex(y)}${hex(w)}${hex(h)}`;

      const values = {
        Thumb: `${this.model.Hash}-${area}`,
        X: x,
        Y: y,
        W: w,
        H: h,
        Q: face.score,
        Invalid: false,
        SubjUID: null,
        Face: face,
      };

      return new Marker(values);
    },
    markerExists(face_marker) {
      const has_similar_marker = (marker) => Util.iou(marker, face_marker) > MarkerOverlapThreshold;
      return this.markers.some(has_similar_marker);
    },
    loadAllFaces() {
      this.model.loadFaces().then((faces) => {
        faces.forEach(face => {
          const face_marker = this.markerFromFace(face);

          if (!this.markerExists(face_marker)) {
            this.markers.push(face_marker);
          }
        });
      });
    },
    onSelect(face) {
      this.dialog.select = false;

      if (this.busy || !face) return;

      if (this.markerExists(this.markerFromFace(face))) {
        Notify.warn(this.$gettext("Face region already exists"));
        return;
      }

      this.busy = true;
      this.$notify.blockUI();

      this.model.addFace(face).then((data) => {
        this.markers.push(new Marker(data));
      }).finally(() => {
        this.$notify.unblockUI();
        this.busy = false;
      });
    },
    onConfirm(marker, index) {
      if (this.busy || !marker) return;

      this.busy = true;
      this.$notify.blockUI();

      this.model.addFace(marker.Face).then((data) => {
        this.$set(this.markers, index, new Marker(data));
      }).finally(() => {
        this.$notify.unblockUI();
        this.busy = false;
      });
    },
    onReject(marker) {
      if (this.busy || !marker) return;

      this.busy = true;
      this.$notify.blockUI();

      marker.reject().finally(() => {
        this.$notify.unblockUI();
        this.busy = false;
      });
    },
    onApprove(marker) {
      if (this.busy || !marker) return;

      this.busy = true;

      marker.approve().finally(() => this.busy = false);
    },
    onClearSubject(marker) {
      if (this.busy || !marker) return;

      this.busy = true;
      this.$notify.blockUI();

      marker.clearSubject(marker).finally(() => {
        this.$notify.unblockUI();
        this.busy = false;
      });
    },
    onRename(marker) {
      if (this.busy || !marker) return;

      this.busy = true;
      this.$notify.blockUI();

      marker.rename().finally(() => {
        this.$notify.unblockUI();
        this.busy = false;
      });
    },
  },
};
</script>
