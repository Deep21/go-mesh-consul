<template>
  <div class="hello">
      <table class="table">
    <thead>
      <tr>
        <th scope="col">Client</th>
        <th scope="col">Email</th>
      </tr>
    </thead>
    <tbody>
      <tr :key="index" v-for="(pays, index) in pays">
        <td>{{ pays.CustumerFirstName }} {{ pays.CustumerLastName }}</td>
        <td>{{ pays.CusterEmail }}</td>
      </tr>
    </tbody>
  </table>
    </div>      
</template>

<script>
import axios from 'axios';

export default {
  name: 'HelloWorld',
  data(){
    return{
      pays:  null,
    }
  },
  mounted() {
    axios.get("http://localhost:9191/v1/orders")
    .then((response) => {
     this.pays = response.data;
      console.log(this.pays )
    });
  },
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
