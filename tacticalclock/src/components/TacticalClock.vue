<template>
<div class="container">
  <div>{{hours}}</div>
  <div>:</div>
  <div>{{minutes}}</div>
  <div>:</div>
  <div>{{seconds}}</div>
</div>
</template>

<script>
export default {
  name: "TacticalClock",
  data () {
    return {
      hours: 0,
      minutes: 0,
      seconds: 0,
    }
  },
  methods: {
    setTime() {
      setInterval(() => {
        const date = new Date()
        this.hours = date.getHours()
        this.minutes = this.checkSingleDigit(date.getMinutes())
        this.seconds = this.checkSingleDigit(date.getSeconds())
      }, 1000)
    },
    checkSingleDigit (digit) {
      return ('0' + digit).slice(-2)
    }
  },
  mounted() {
    this.setTime()
  }
}
</script>

<style scoped>

</style>