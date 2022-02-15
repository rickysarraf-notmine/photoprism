<template>
  <v-dialog v-model="show" lazy persistent max-width="1000" class="p-people-select-face-dialog" @keydown.esc="cancel">
    <v-card raised elevation="24">
      <v-container fluid class="pb-2 pr-2 pl-2 pt-2">
        <v-layout row wrap>
          <v-flex text-xs-left align-self-center>
            <div class="subheading pb-2">
                <translate>Select face area</translate>
            </div>
            <cropper
              class="cropper"
              :src="file.getDownloadUrl()"
              :stencil-props="{aspectRatio: 1}"
              :defaultSize="{width: 250, height: 250}"
              @change="onChange"
            ></cropper>
          </v-flex>
          <v-flex text-xs-right class="pt-3">
            <v-btn depressed color="secondary-light" class="action-cancel" @click.stop="cancel">
              <translate key="Cancel">Cancel</translate>
            </v-btn>
            <v-btn color="primary-button" depressed dark class="action-confirm"
                   @click.stop="confirm">
              <translate key="Select">Select</translate>
            </v-btn>
          </v-flex>
        </v-layout>
      </v-container>
    </v-card>
  </v-dialog>
</template>
<script>
import 'vue-advanced-cropper/dist/style.css';
import File from "model/file";

export default {
  name: 'PPeopleSelectFaceDialog',
  props: {
    show: Boolean,
    file: File,
  },
  data() {
    return {
      coordinates: null,
    };
  },
  methods: {
    onChange({coordinates}) {
      this.coordinates = coordinates;
    },
    cancel() {
      this.$emit('cancel');
    },
    confirm() {
      const size = this.coordinates.width;
      const half_size = parseInt(size / 2);

      const face = {
        cols: this.file.Width,
        rows: this.file.Height,
        score: 0,
        face: {
          name: "face",
          size: size,
          x: this.coordinates.top + half_size,
          y: this.coordinates.left + half_size,
        },
      };

      this.$emit('confirm', face);
    },
  }
};
</script>
