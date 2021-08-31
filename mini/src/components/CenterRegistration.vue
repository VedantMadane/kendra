<template>
    <div>
        <form action="/" @submit.prevent="sended()"> 
      <!-- <div class="row mt2">
        <div class="col-md-6">
          
        </div>
      </div> -->
    
      <div class="container p-5">
        <div class="card shadow">
          <div class="card-header">Center Registration</div>
          <div class="card-body">
            <div class="row g-1">
                <div class="col-md-6">
                  <label for="centername">Center Name</label>
                  <input class="form-control" type="text" name="centername" ref="centername" id="centername" placeholder="Center Name">
                </div>
                <div class="col-md-6">
                  <label for="ownername">Owner Name</label>
                  <input class="form-control" type="text" name="ownername" ref="ownername" id="ownername" placeholder="Owner Name">
                </div>
            <div class="row">
              <div class="col-md-6">
                <label for="ownermobilenumber">Owner Mobile Number</label>
                <input class="form-control" type="number" name="ownermobilenumber" ref="ownermobilenumber" id="ownermobilenumber" placeholder="Owner Mobile Number">
              </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label for="owneremailid">Owner Email ID</label>
                  <input class="form-control" type="email" name="owneremailid" ref="owneremailid" id="owneremailid" placeholder="Owner Email ID">
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <label for="lf">Learning Facilitator Name (LF)</label>
                  <input class="form-control" type="text" name="lf" ref="lf" id="lf" placeholder="Learning Facilitator Name (LF)">
                </div>
                <div class="col-md-6">
                  <label for="lfemailid">Learning Facilitator Email ID</label>
                  <input class="form-control" type="email" name="lfemailid" ref="lfemailid" id="lfemailid" placeholder="Learning Facilitator Email ID">
                </div>
              </div>
              <div class="row">
                

              <div class="col-md-6">
                <label for="lfmobilenumber">Learning Facilitator Mobile Number</label>
                <input class="form-control" type="number" name="lfmobilenumber" ref="lfmobilenumber" id="lfmobilenumber" placeholder="Learning Facilitator Mobile Number">
              </div>
            </div>
            <div class="row p-3">
                <div class="col-md-6">
                  <label for="city">City</label>
                  <input class="form-control" type="text" name="city" ref="city" id="city" placeholder="City">
</div>
            <div class="row p-3">
                <div class="col-md-6">
                  <label for="pincode">Pin Code</label>
                  <input class="form-control" type="number" name="pincode" ref="pincode" id="pincode" placeholder="Pin Code">
                  
                </div>
            </div>
            
                <div class="col-md-12">
                  <label for="address">Address</label>
                  <input class="form-control" type="textarea" name="address" ref="address" id="address" placeholder="Address">
                </div>
            </div>
            <div class="row">
<div class="row p-3">
            <div>
                <div class="col-md-6">
                  <label for="password">Password</label>
                  <input class="form-control" type="password" pattern=".{8,}" name="password" ref="password" id="password" placeholder="Password">
                </div>
            </div>

            
                <div class="col-md-6">
                  <label for="confirmpassword">Confirm Password</label>
                  <input class="form-control" type="password" pattern = ".{8,}" name="confirmpassword" ref="confirmpassword" id="confirmpassword" placeholder="Confirm password">
                </div>
  </div>
            </div>

            <div class="row">
              <div class="col-md-6">
                
              </div>
              <div class="col-md-6">
                <button class="btn btn-outline-primary mr-3">Create Center</button> 
                <!--&nbsp-->
                <input type="reset" class="btn btn-outline-danger" name="" id="">
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    </form>
    </div>
</template>
<script>
import axios from 'axios'
export default {
    name:"CenterRegistration",
    methods:{
    sended(){
      var centername = this.$refs.centername.value
      var ownername = this.$refs.ownername.value
      var ownermobilenumber = this.$refs.ownermobilenumber.value
      var owneremailid = this.$refs.owneremailid.value
      var lf = this.$refs.lf.value
      var lfemailid = this.$refs.lfemailid.value
      var lfmobilenumber = this.$refs.lfmobilenumber.value
      var city = this.$refs.city.value
      var pincode = this.$refs.pincode.value
      var address = this.$refs.address.value
      var password = this.$refs.password.value
      var confirmpassword = this.$refs.confirmpassword.value
      //var q = document.getElementById("gender").value
      console.log(centername,ownername,ownermobilenumber, ownermobilenumber,owneremailid,lf,lfemailid,lfmobilenumber,city, pincode, address, password, confirmpassword)
      axios.post('http://localhost:3000/data',{centername:centername, ownername:ownername, ownermobilenumber:ownermobilenumber, owneremailid:owneremailid, lf:lf, lfemailid:lfemailid, 
      lfmobilenumber:lfmobilenumber, city:city, pincode:pincode, address:address, password:password, confirmpassword:confirmpassword})
      .then(res => {
            console.log(res.data);
        })
    }
  },
  
}
</script>

<!--
<script>
          
            // Function to check Whether both passwords
            // is same or not.
            function checkPassword(form) {
                password1 = form.password1.value;
                password2 = form.password2.value;
  
                // If password not entered
                if (password1 == '')
                    alert ("Please enter Password");
                      
                // If confirm password not entered
                else if (password2 == '')
                    alert ("Please enter confirm password");
                      
                // If Not same return False.    
                else if (password1 != password2) {
                    alert ("\nPassword did not match: Please try again...")
                    return false;
                }
  
                // If same return True.
                else{
                    alert("Password Match: Welcome to GeeksforGeeks!")
                    return true;
                }
            }
        </script>

<script>

var password = document.getElementById("password")
  , confirmpassword = document.getElementById("confirmpassword");

function validatePassword(){
  if(password.value != confirmpassword.value) {
    confirmpassword.setCustomValidity("Passwords Don't Match");
  } else {
    confirmpassword.setCustomValidity('');
  }
}

password.onchange = validatePassword;
confirmpassword.onkeyup = validatePassword;
</script>


<script>  
function verifyPassword() {  
  var pw = document.getElementById("password").value;  
  //check empty password field  
  if(pw == "") {  
     document.getElementById("message").innerHTML = "**Fill the password please!";  
     return false;  
  }  
   
 //minimum password length validation  
  if(pw.length < 8) {  
     document.getElementById("message").innerHTML = "**Password length must be atleast 8 characters";  
     return false;  
  }  
  
//maximum length of password validation  
  if(pw.length > 15) {  
     document.getElementById("message").innerHTML = "**Password length must not exceed 15 characters";  
     return false;  
  } else {  
     alert("Password is correct");  
  }  
}  
</script>  

-->