<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr>
        <form-tag @myevent="submitHandelr">

          <text-input
            v-model="email"
            label="Email"
            name="email"
            type="email"
            required="true">
          </text-input>

          <text-input
            v-model="password"
            label="Password"
            name="password"
            type="password"
            required="true">
          </text-input>

          <hr>
          <input type="submit" class="btn btn-primary" value="Login">

        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>
import FormTag from './forms/FormTag.vue'
import TextInput from './forms/TextInput.vue'

export default{
  name: 'login', 
  components: {
    FormTag,
    TextInput,
  },
  data()  {
    return {
      email: "",
      password: "",
    }
  }
  methods: {
    submitHandelr() {
      console.log("submitHandelr called - success");
    }
  },
  mounted() {
      (function () {
        'use strict'

        // Fetch all the forms we want to apply custom Bootstrap validation styles to
        var forms = document.querySelectorAll('.needs-validation')

        // Loop over them and prevent submission
        Array.prototype.slice.call(forms)
            .forEach(function (form) {
            form.addEventListener('submit', function (event) {
                if (!form.checkValidity()) {
                event.preventDefault()
                event.stopPropagation()
                }

                form.classList.add('was-validated')
            }, false)
            })
        })()
  }
}
</script>