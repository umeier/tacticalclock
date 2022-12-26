<template>
  <div>
    <TacticalClock/>
  </div>

</template>

<script>
import TacticalClock from "@/components/TacticalClock.vue";

export default {
  name: "App",
  components: {TacticalClock}
}
</script>

<style lang="scss">
@import "./assets/main.scss";
</style>