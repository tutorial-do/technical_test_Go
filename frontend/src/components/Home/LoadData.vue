<template>
  <section class="section">
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <v-card
          elevation="5"
          >
          <h1>LOAD DATA</h1>
          </v-card>
          <template>
            <v-menu
              ref="menu"
              v-model="menu"
              :close-on-content-click="false"
              transition="scale-transition"
              offset-y
              min-width="290px"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-text-field
                  v-model="date"
                  label="Select date to load data"
                  prepend-icon="mdi-calendar"
                  readonly
                  v-bind="attrs"
                  v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker
                ref="picker"
                v-model="date"
                :max="new Date().toISOString().substr(0, 10)"
                min="1950-01-01"
                @change="save"
              ></v-date-picker>
            </v-menu>
          </template>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
export default {
  name: 'LoadData',
  // data() {
  //   return {
  //     date: null,
  //     menu: false,
  //   }
  // },
  // watch: {
  //   menu(val) {
  //     val && setTimeout(() => (this.$refs.picker.activePicker = 'YEAR'));
  //   },
  // },
  // methods: {
  //   save(date) {
  //     this.$refs.menu.save(date);
  //   },
  // },
};
</script>
<style scoped lang="scss">
</style>
