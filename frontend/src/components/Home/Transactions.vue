<template>
  <v-card
  elevation="5"
  >
    <v-card-text>
      <h1>Transactions of Buyer: {{ activeBuyer.name }}</h1>
    </v-card-text>
    <v-data-table
      :headers="headers"
      :items="transactionsByBuyerID"
      :items-per-page="5"
      class="elevation-1"
    >
    <template
    v-slot:item.select="{ item }"
    >
      <v-icon
        small
        class="mr-2"
        @click="setterTransaction(item.id)
                setterDevice(item.ip)"
      >
        Select to display
      </v-icon>
    </template>
    </v-data-table>
  </v-card>
</template>
<script>

import { mapState } from 'vuex';

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
    return {
      headers: [
        {
          text: 'Transaction ID',
          align: 'start',
          sortable: false,
          value: 'id',
        },
        { text: 'Device', value: 'device' },
        { text: 'IP', value: 'ip' },
        { text: 'Buyer ID', value: 'buyerID' },
        { text: 'Selector', value: 'select', sortable: false },
      ],
    };
  },
  computed: mapState({
    transactionsByBuyerID: (state) => state.transactionsByBuyerID,
    activeBuyer: (state) => state.activeBuyer,
  }),
  mounted() {},
  methods: {
    setterTransaction(currentTrx) {
      this.$store.dispatch('productsByTransactionID', currentTrx);
    },
    setterDevice(ipAddress) {
      this.$store.dispatch('getBuyersByIP', ipAddress);
    },
  },
};
</script>
<style scoped lang="scss">
</style>
