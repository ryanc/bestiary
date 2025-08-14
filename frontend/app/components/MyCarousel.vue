<script setup lang="ts">
const runtimeConfig = useRuntimeConfig();

let baseUrl;

if (import.meta.server) {
  baseUrl = runtimeConfig.apiBase;
} else {
  baseUrl = runtimeConfig.public.apiBase;
}

const { data, pending, error, refresh } = await useFetch<string[]>(`${baseUrl}/beasts`);

const items = computed(() => data.value || [])
</script>

<template>
  <div>
    <div v-if="pending" class="flex justify-center items-center h-96">
      <div class="text-center">
        <UIcon name="i-heroicons-photo" class="w-12 h-12 text-gray-400 mx-auto mb-4 animate-pulse" />
        <p class="text-gray-600">Loading photos...</p>
      </div>
    </div>
    <UAlert
      v-else-if="error"
      icon="i-heroicons-exclamation-triangle"
      color="error"
      variant="soft"
      title="Failed to load photos"
      :description="error.message"
      class="mb-6"
    >
      <template #actions>
        <UButton @click="refresh()" variant="ghost" size="xs">
          Try Again
        </UButton>
      </template>
    </UAlert>
    <UCarousel
      v-else
      v-slot="{ item }"
      auto-height
      arrows
      :autoplay="{ delay: 5000 }"
      :ui="{ item: 'basis-full' }"
      :items="items"
      class="w-full max-w-3xl mx-auto"
    >
      <img
        :src="item"
        width="1000"
        height="1000"
        class="rounded-lg"
      >
    </UCarousel>
  </div>
</template>
