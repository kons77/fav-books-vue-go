<template>
    <form 
      @submit.prevent="submit"
      :ref="name"
      :event="event"
      autocomplete="off" 
      :method="method" 
      :action="action" 
      class="needs-validation" 
      novalidate>
      <slot></slot>
    </form>
</template>

<script>
export default{
  name: 'FormTag',
  props: ["method", "action", "name", "event"],
  methods: {
    submit() {
      let myForm = this.$refs[this.$props.name];

      // standart bootstrap validation but simplified
      if (myForm.checkValidity()) {
        console.log("My event name is", this.$props['event']);
        console.log("Name", this.$props.name);
        this.$emit(this.$props['event']); // emitting an event here instead of submitting the form
      }
      myForm.classList.add('was-validated');
    }
  },
}

</script>