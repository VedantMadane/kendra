<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <MyHeader/>
    <CenterRegistration/>
    <!-- <form action="/" @submit.prevent="pathav()">  -->
      <!-- <div class="row mt2">
        <div class="col-md-6">
          
        </div>
      </div>
      <div class="container p-5">
        <div class="card shadow">
          <div class="card-header">Register the student</div>
          <div class="card-body">
            <div class="row">
                <div class="col-md-6">
                  <label for="fullname">Name</label>
                  <input class="form-control" type="text" name="fullname" ref="fullname" id="name" placeholder="Full Name">
                </div>
                <div class="col-md-6">
                  <label for="mothername">MotherName</label>
                  <input class="form-control" type="text" name="mothername" ref="mothername" id="mothername" placeholder="Mothers Name">
                </div>
            </div>
            <div class="row p-3">
              <div class="col-md-6">
                  <label for="date">date</label>
                  <input class="form-control" type="date" name="date" ref="date" id="date" placeholder="Full Name">
              </div>
              <div class="col-md-6">
                <label for="nationality">Natioanlity</label>
                <input class="form-control" type="text" name="nationality" ref="nationality" id="nationalit" placeholder="nationality">
              </div>
            </div>
            <div class="row p-3">
                <div class="col-md-6">
                  <label for="religion">Religion</label>
                  <input class="form-control" type="text" name="religion" ref="religion" id="religion" placeholder="Religion">
                </div>
                <div class="col-md-6">
                  <label for="caste">Caste</label>
                  <input class="form-control" type="text" name="caste" ref="caste" id="caste" placeholder="Caste">
                </div>
            </div>

            <div class="row p-3">
                <div class="col-md-6">
                  <label for="gender">Gender</label>
                  <input class="form-control" type="text" name="gender" ref="gender" id="gender" placeholder="Gender">
                  
                </div>
                <div class="col-md-6">
                  <label for="school">School Name</label>
                  <input class="form-control" type="text" name="school" ref="school" id="school" placeholder="Previous School Name">
                </div>
            </div>

            <div class="row p-3">
                <div class="col-md-6">
                  <label for="result">Result</label>
                  <input class="form-control" type="text" name="result" ref="result" id="result" placeholder="Result">
                  
                </div>
                <div class="col-md-6">
                  <label for="remark">Remark</label>
                  <input class="form-control" type="remark" name="remark" ref="remark" id="remark" placeholder="Remarks">
                </div>
            </div>

            <div class="row p-3">
                <div class="col-md-6">
                  <label for="image">Result</label>
                  <input class="form-control" type="file" name="image" ref="image" id="image" >
                  
                </div>
  
            </div>

            <div class="row ">
              <div class="col-md-6">
                
              </div>
              <div class="col-md-6">
                <button class="btn btn-outline-primary">Submit Name</button>&nbsp
                <input type="reset" class="btn btn-outline-danger" name="" id="">
              </div>
            </div>
          </div>
        </div>
      </div>
      
    </form> -->
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import MyHeader from '@/components/MyHeader.vue'
import CenterRegistration from '@/components/CenterRegistration.vue'
import axios from 'axios'
export default {
  name: 'Home',
  components: {
    HelloWorld,
    MyHeader,
    CenterRegistration,
  },
  // methods:{
  //   pathav(){
  //     var fullname = this.$refs.fullname.value
  //     var date = this.$refs.date.value
  //     var mothername = this.$refs.mothername.value
  //     var nationality = this.$refs.nationality.value
  //     var religion = this.$refs.religion.value
  //     var caste = this.$refs.caste.value
  //     var gender = this.$refs.gender.value
  //     var school = this.$refs.school.value
  //     var result = this.$refs.result.value
  //     var remark = this.$refs.remark.value
  //     var image = this.$refs.image.value
  //     //var q = document.getElementById("gender").value
  //     console.log(fullname,date,mothername,nationality,religion,caste,gender,school,result,remark,image)
  //     axios.post('http://localhost:3000/data',{fullname:fullname, date:date, mothername:mothername, nationality:nationality, religion:religion, caste:caste, gender:gender, school:school, result:result, remark:remark, image:image })
  //     .then(res => {
  //           console.log(res.data);
  //       })
  //   }
  // }
}
</script>
