<template>
  <div class="container">
    <div class="tclock">
      <div>{{ days }}<span class="time">{{ hours }}{{ minutes }}</span>{{ months }}{{ years }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "TacticalClock",
  data() {
    return {
      days: 0,
      hours: 0,
      minutes: 0,
      months: "",
      years: ""
    }
  },
  methods: {
    setTime() {
      setInterval(() => {
        const date = new Date()
        this.days = date.getDate()
        this.hours = date.getHours()
        this.minutes = this.checkSingleDigit(date.getMinutes())
        this.months = this.getMonthText(date.getMonth())
        this.years = date.getFullYear().toString().slice(-2)
      }, 1000)
    },
    checkSingleDigit(digit) {
      return ('0' + digit).slice(-2)
    },
    getMonthText(digit) {
      const month = ["jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"]
      return month[digit]
    }
  },
  mounted() {
    this.setTime()
  }
}
</script>

<style scoped>
@font-face {
  font-family: 'MyLubalin';
  src: url("/fonts/LubalinGraphStd/LubalinGraphStd-Demi.woff") format('woff'),
  url("/fonts/LubalinGraphStd/LubalinGraphStd-Demi.woff2") format('woff2');
}

.tclock {
  position: fixed; /* or absolute */
  top: 50%;
  left: 50%;
  /* bring your own prefixes */
  transform: translate(-50%, -50%);
  font-size: 10vw;
  font-family: MyLubalin, sans-serif;
}
.time {
  font-size: larger;
}
</style>