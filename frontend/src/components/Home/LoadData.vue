<template>
  <section class="section">
    <div class="container">
      <div class="row">
        <div class="col-md-6">
          <v-card
          elevation="5"
          >
          <h1>LOAD DATA</h1>
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
              v-on:click="loadData(picker)"
              >
                Load data
              </v-btn>
            </template>
          </v-card>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
export default {
  name: 'LoadData',
  data() {
    return {
      picker: new Date().toISOString().substr(0, 10),
    };
  },
  methods: {
    loadData(picker) {
      this.$store.dispatch('loadData', picker);
    },
    deleteData() {
      this.$store.dispatch('deleteData');
    },
  },
};
</script>
<style scoped lang="scss">
</style>
