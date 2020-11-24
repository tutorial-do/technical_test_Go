<template>
  <v-card
  elevation="5"
  >
  <v-card-text>
    <h1>Load Data</h1>
  </v-card-text>
    <template>
      <v-row justify="center">
        <v-date-picker
          v-model="picker"
          year-icon="mdi-calendar-blank"
          prev-icon="mdi-skip-previous"
          next-icon="mdi-skip-next"
          color="light-blue"
          :max="new Date().toISOString().substr(0, 10)"
        ></v-date-picker>
      </v-row>
    </template>
    <template>
      <v-btn
      block
      :loading=loadingStatus
      :disabled=loadingStatus
      color="light-blue"
      class="white--text"
      v-on:click="loadData(picker)"
      >
        Load data
      </v-btn>
      <v-btn
      block
      v-on:click="deleteData()"
      >
        Delete all data
      </v-btn>
    </template>
  </v-card>
</template>
<script>
import { mapState } from 'vuex';

export default {
  name: 'LoadData',
  data() {
    return {
      picker: new Date().toISOString().substr(0, 10),
    };
  },
  computed: mapState({
    loadingStatus: (state) => state.loadingStatus,
  }),
  methods: {
    loadData(picker) {
      this.$store.dispatch('loadData', picker);
    },
    deleteData() {
      this.$store.dispatch('deleteData');
    },
  },
  // watch: {
  //   loader () {
  //     const l = this.loader
  //     this[l] = !this[l]

  //     setTimeout(() => (this[l] = false), 3000)

  //     this.loader = null
  //   },
  // },
};
</script>
<style scoped lang="scss">
</style>
