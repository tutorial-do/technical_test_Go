<template>
  <div class="container container--fluid">
      <div class="row">
          <div class="col-md-6">
            <AllBuyers
            :data="allBuyers"
            />
          </div>
          <div class="col-md-6">
            <BuyerInfo
            :data="activeBuyer"
            />
            <LoadData />
          </div>
      </div>
      <div class="row">
          <div class="col-md-6">
            <Transactions
            :data="allTransactions"
            />
          </div>
          <div class="col-3">
            <ProductsByTransaction />
          </div>
          <div class="col-3">
            <BuyersByIP />
          </div>
      </div>
      <div class="row">
          <div class="col-12">
            <Products
            :data="allProducts"
            />
          </div>
      </div>
  </div>
</template>

<script>
import AllBuyers from '@/components/Home/AllBuyers.vue';
import BuyerInfo from '@/components/Home/BuyerInfo.vue';
import BuyersByIP from '@/components/Home/BuyersByIP.vue';
import Transactions from '@/components/Home/Transactions.vue';
import Products from '@/components/Home/Products.vue';
import ProductsByTransaction from '@/components/Home/ProductsByTrans.vue';
import { mapState } from 'vuex';
import LoadData from '../components/Home/LoadData.vue';

export default {
  name: 'Home',
  components: {
    AllBuyers,
    BuyerInfo,
    BuyersByIP,
    LoadData,
    Products,
    ProductsByTransaction,
    Transactions,
  },
  data() {
    return {};
  },
  created() {
    this.$store.dispatch('fetchAllBuyers');
    // this.$store.dispatch('fetchAllProducts');
    // this.$store.dispatch('fetchAllTransactions');
  },
  computed: mapState({
    allBuyers: (state) => state.allBuyers,
    allProducts: (state) => state.allProducts,
    allTransactions: (state) => state.allTransactions,
    activeBuyer: (state) => state.activeBuyer,
  }),
};

</script>
<style scoped lang="scss">
</style>
