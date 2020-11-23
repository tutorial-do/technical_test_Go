<template>
  <v-card>
    <v-card-text>
      <h1>All Buyers</h1>
    </v-card-text>
    <v-data-table
      :headers="headers"
      :items="data"
      :items-per-page="5"
    >
    <template v-slot:item.select="{ item }">
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
