<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{label}}</label>
    <select 
      :id="name" 
      class="form-select" 
      :name="name" 
      :required="required" 
      :multiple="multiple"
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)">
      
      <option disabled value="">Choose...</option> <!-- no ability to choose on iPhone without this line -->
      <option v-for="option in items" :value="option.value" :key="option.value">
          {{option.text}}
      </option>
    </select>
  </div>
</template>


<script>
export default{
  name: 'SelectInput',
  props: ["items", "name", "required", "label", "modelValue", "multiple"],
  emits: ['update:modelValue'], // don't have to but it's a good practice
}
</script>

