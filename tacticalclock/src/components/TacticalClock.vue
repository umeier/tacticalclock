<template>
  <div class="container">
    <div class="tclock">
      <div v-if="years">{{ days }}<span class="time">{{ hours }}{{ minutes }}</span><span class="seconds">{{ seconds }}</span><span class="tz">{{ timezone }}</span>{{ months }}{{ years }}</div>
      <div v-else>Initializing</div>
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
      seconds: 0,
      months: "",
      years: "",
      timezone: ""
    }
  },
  methods: {
    setTime() {
      setInterval(() => {
        const date = new Date()
        this.days = this.checkSingleDigit(date.getDate())
        this.hours = this.checkSingleDigit(date.getHours())
        this.minutes = this.checkSingleDigit(date.getMinutes())
        this.seconds = this.checkSingleDigit(date.getSeconds())
        this.months = this.getMonthText(date.getMonth())
        this.years = date.getFullYear().toString().slice(-2)
        this.timezone = this.getTimeZone(date.getTimezoneOffset())
      }, 1000)
    },
    checkSingleDigit(digit) {
      return ('0' + digit).slice(-2)
    },
    getMonthText(digit) {
      return this.month[digit]
    },
    getTimeZone(number) {
      let index = 12 + (number / 60 * -1)
      return this.zones[index]
    },
  },

  mounted() {
    this.setTime()
  },
  created() {
    this.month = ["jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"]
    this.zones = ["Y", "X", "W", "V", "U", "T", "S", "R", "Q", "P", "O", "N", "Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "K", "L", "M"]
  }
}
</script>

<style scoped>
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
.seconds {
  font-size: smaller;
  color: gray;
}
.tz {
  color: gray;
}
</style>