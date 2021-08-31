<template>
    <div>
        <table>
      <tr>
        <th>Student Id</th>
        <th>Student Name</th>
        <th>Total Subjects</th>
      </tr>
      <tr  v-for="(hello,i) in Students.data" :key="i">
        <td>{{hello.studId}}</td>
        <td>{{hello.studName}}</td>
        <td>{{hello.subjectOpted}}</td>
      </tr>
    </table>
    </div>
</template>

<script>
import axios from 'axios'
export default {
    name:'Delete',
    data:{
         Students:[],
    },
    created(){
         let amol=this
        axios.get("http://localhost:3000/delete")
     
    
  .then(function (response) {
    // handle success
    amol.Students=response.data
    
  })
  .catch(function (error) {
    // handle error
    console.log(error);
  })
  .then(function () {
    // always executed
  });  
    }
}
</script>