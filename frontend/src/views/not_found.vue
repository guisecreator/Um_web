<template>
  <div class="fullscreen text-white text-center q-pa-md flex flex-center">
    <div>
      <div style="font-size: 30vh" color="dark">
        404
      </div>
      <div class="text-h2" style="opacity:.4">
        Error. Page not found.
      </div>
      <v-card-title class="text-center my-5">You've come to the non-existent page.</v-card-title>
      <v-btn
        class="q-mt-xl"
        color="white"
        text-color="blue"
        unelevated
        to="/"
        label="Go Home"
        no-caps
      >Login
      </v-btn>
    </div>
  </div>
</template>

<style scoped>
.v-card-title {
  margin-top: 1rem;
}

.v-btn {
  margin-top: 1rem;
}

.my-5 {
    margin-top: 59px !important;
    margin-bottom: 28px !important;
}

.v-btn[data-v-2f8da421] {
    margin-top: -3rem;
}
</style>

<script setup lang="ts">
</script>
