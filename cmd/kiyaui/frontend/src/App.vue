<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" :clipped="$vuetify.breakpoint.lgAndUp"
                         app>
      <profile-list></profile-list>
    </v-navigation-drawer>
    <v-app-bar app fixed clipped-left>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title
              style="width: 300px"
              class="ml-0 pl-4"
      >
        <span class="hidden-sm-and-down">Kiya UI</span>
      </v-toolbar-title>
    </v-app-bar>
    <v-content>
      <v-container fluid class="px-0">
      <v-layout fill-height class="px-0">
        <profile-details></profile-details>
      </v-layout>
      </v-container>
    </v-content>
    <v-footer app fixed>
      <span style="margin-left:1em">&copy; You</span>
    </v-footer>
  </v-app>
</template>

<script>
  import ProfileList from "./components/ProfileList.vue"
  import ProfileDetails from "./components/ProfileDetails.vue"

  export default {
    data: () => ({
      drawer: false
    }),
    components: {
      ProfileList,
      ProfileDetails,
    },
    props: {
      source: String
    }
  }
</script>

<style>
</style>
