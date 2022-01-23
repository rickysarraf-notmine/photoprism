<template>
  <div v-if="show" class="sphere-viewer" role="dialog" @keydown.esc.stop.prevent="onClose">
      <v-pannellum v-show="show" ref="viewer" :src="source" :showZoom=showZoom :showFullscreen=showFullscreen :hfov=hfov></v-pannellum>
  </div>
</template>
<script>
import Event from "pubsub-js";

export default {
  name: 'PSphereViewer',
  data() {
    return {
      show: false,
      source: "",
      showZoom: true,
      showFullscreen: true,
      hfov: 120,
      subscriptions: [],
    };
  },
  created() {
    this.subscriptions['sphereviewer.open'] = Event.subscribe('sphereviewer.open', this.onOpen);
    this.subscriptions['sphereviewer.close'] = Event.subscribe('sphereviewer.close', this.onClose);
  },
  beforeDestroy() {
    for (let i = 0; i < this.subscriptions.length; i++) {
      Event.unsubscribe(this.subscriptions[i]);
    }

    this.onClose();
  },
  methods: {
    onOpen(ev, item) {
      this.show = true;
      this.source = item.download_url;

      this.$nextTick(() => this.$refs.viewer.$el.focus());
    },
    onClose() {
      this.show = false;
    },
  },
};
</script>
