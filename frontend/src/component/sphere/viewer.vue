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
  beforeMount() {
    window.addEventListener('hashchange', this.onClose, {passive: true});
  },
  beforeDestroy() {
    window.removeEventListener('hashchange', this.onClose, false);

    for (let i = 0; i < this.subscriptions.length; i++) {
      Event.unsubscribe(this.subscriptions[i]);
    }

    this.onClose();
  },
  methods: {
    onOpen(ev, item) {
      this.show = true;
      this.source = item.DownloadUrl;

      // Enable closing the photosphere viewer on mobile using the back button by creating a new
      // history state, which will be poped from the history stack when the back button is pressed.
      // In turn this will trigger the "hashchange" event, which will close the viewer.
      history.pushState("", "", "#sphereviewer");

      this.$nextTick(() => this.$refs.viewer.$el.focus());
    },
    onClose() {
      this.show = false;
    },
  },
};
</script>
