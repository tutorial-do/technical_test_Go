<template>
  <div class="container container--fluid">
      <div class="row">
          <div class="col-md-4">
            <AllBuyers
            :data="allBuyers"
            />
          </div>
          <div class="col-md-4">
            <LoadData />
          </div>
          <div class="col-md-4">
            <Transactions />
          </div>
      </div>
      <div class="row">
          <div class="col-12">
            <Products />
          </div>
      </div>
  </div>
</template>

<script>
import AllBuyers from '@/components/Home/AllBuyers.vue';
import Transactions from '@/components/Home/Transactions.vue';
import Products from '@/components/Home/Products.vue';
import { mapState } from 'vuex';
import LoadData from '../components/Home/LoadData.vue';

export default {
  name: 'Home',
  components: {
    AllBuyers,
    LoadData,
    Products,
    Transactions,
  },
  data() {
    return {};
  },
  created() {
    this.$store.dispatch('fetchAllBuyers');
  },
  computed: mapState({
    allBuyers: (state) => state.allBuyers,
  }),
};

</script>
<style scoped lang="scss">
</style>
