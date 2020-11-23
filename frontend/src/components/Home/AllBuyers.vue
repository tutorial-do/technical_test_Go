<template>
  <section class="section">
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <v-card
            elevation="5"
          >
          <h1>ALL BUYERS</h1>
          <v-data-table
            :headers="headers"
            :items="data"
            :items-per-page="5"
            class="elevation-1"
          >
          <template v-slot:item.select="{ item }">
          <!-- <template v-slot:[`item.select`]> -->
            <v-icon
              small
              class="mr-2"
              @click="setterBuyer(item.name, item.id, item.age)"
            >
              Select to display
            </v-icon>
          </template>
          </v-data-table>
          </v-card>
        </div>
      </div>
    </div>
  </section>
</template>
<script>

export default {
  name: 'BuyersByIP',
  props: {
    data: {
      type: [Array, String],
      default: '',
      required: false,
    },
  },
  data() {
    return {
      headers: [
        {
          text: 'Name',
          align: 'start',
          sortable: false,
          value: 'name',
        },
        { text: 'ID', value: 'id' },
        { text: 'Age', value: 'age' },
        { text: 'Selector', value: 'select', sortable: false },
      ],
    };
  },
  methods: {
    setterBuyer(BuyerName, BuyerID, BuyerAge) {
      this.$store.dispatch('setCurrentBuyer', {
        name: BuyerName,
        id: BuyerID,
        age: BuyerAge,
      });
    },
  },
};
</script>
<style scoped lang="scss">
</style>
