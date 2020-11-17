<template>
  <section class="section">
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <v-card
          elevation="5"
          >
          <h1>TRANSACTIONS</h1>
          <ul class="sidebar-group-items">
            <li
              v-for="(item, index) in data"
              :key="index"
            >
              {{ item.buyerID }}
              {{ item.device }}
              {{ item.ProductsIDs }}
            </li>
          </ul>
          </v-card>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
export default {
  name: 'Transactions',
  props: {
    data: {
      type: [Array, String],
      default: '',
      required: false,
    },
  },
  data() {
    return {};
  },
  mounted() {},
  methods: {},
};
</script>
<style scoped lang="scss">
</style>
