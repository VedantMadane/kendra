<template>
    <div>
        <div class="container-fluid">
            <nav class="navbar navbar-expand-lg navbar-light bg-light sticky-top">
                <div>
                    <span id="navBrand">MS-CIT Center Registration</span>
                </div>
                <div id="nav">
                    <router-link id="hello" to="/">Home</router-link>
                    <router-link id="hello" to="/about">About</router-link>
                    <router-link id="hello" to="/login">login</router-link>
                </div>
            </nav>
        </div>
    </div>
</template>
<style scoped>
    #navBrand{
        margin-left: 1.5rem;
        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
        font-size: 22px;
        font-weight: 500;
        color: #42b983;
    }
    .navbar{
        display: flex;
        justify-content: space-between;
    }
</style>

<script>
export default{
    name:'MyHeader',
}
</script>