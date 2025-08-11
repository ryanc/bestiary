<script setup lang="ts">
const { data } = await useFetch("/api/beasts", {
  lazy: true,
  server: false,
});
</script>

<template>
  <div v-if="data == undefined">
    No data
  </div>
  <UCarousel
    v-slot="{ item }"
    auto-height
    arrows
    :items="data"
    :autoplay="{ delay: 5000 }"
    :ui="{ item: 'basis-full' }" 
    class="w-full max-w-3xl mx-auto"
  >

    <img
      :src="item"
      width="1000"
      height="1000"
      class="rounded-lg"
    >

  </UCarousel>
</template>
