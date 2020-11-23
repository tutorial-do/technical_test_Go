<template>
  <v-card
  elevation="5"
  >
    <v-card-text>
      <h1>Recomended Products</h1>
    </v-card-text>
    <v-data-table
      :headers="headers"
      :items="data"
      :items-per-page="5"
      class="elevation-1"
    ></v-data-table>
  </v-card>
</template>
<script>
export default {
  name: 'Products',
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
          text: 'Product Name',
          align: 'start',
          sortable: false,
          value: 'name',
        },
        { text: 'Product ID', value: 'id' },
        { text: 'Product Price $', value: 'price' },
      ],
    };
  },
  mounted() {},
  methods: {},
};
</script>
<style scoped lang="scss">
</style>
