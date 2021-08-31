<template>
  <div class="about">
    <MyHeader/>
    <Delete/>
  </div>
</template>

<script>
import MyHeader from '@/components/MyHeader.vue'
import Delete from '@/components/Delete.vue'

export default {
  name:"About",
  components:{
    MyHeader,
    Delete,
  }
}
</script>

