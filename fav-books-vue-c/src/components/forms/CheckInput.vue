<!-- converted to the Composition API -->
<template>
    <div class="form-check">
      <input class="form-check-input" 
      :checked="modelValue" 
      :required="required" 
      type="checkbox" 
      :value="value" 
      :name="name" 
      :id="name"
      @input="emit('update:modelValue', $event.target.value)"  >
      <label class="form-check-label" :for="name">
          {{label}}
      </label>
    </div>
</template>


<script setup>
  // export default { name: 'CheckInput', }
  const props = defineProps({
    label: String, 
    name: String, 
    required: Boolean,
    value: String, 
    modelValue: Boolean, 
  })

  const emit = defineEmits(['update:modelValue'])
</script>