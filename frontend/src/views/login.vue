<template>
<auth/>
</template>

<script lang="ts" setup>
import auth from '@/components/auth/auth.vue';
</script>

<style>
.text-center {
  text-align: center;
}
.custom-input .v-input__control {
  height: 55px;
  width: 100%;
}
.custom-input .v-input__input {
  font-size: 14px;
}

.error-button {
  background-color: red !important;
  color: white !important;
}

.success-button {
  background-color: green !important;
  color: white !important;
}
</style>
