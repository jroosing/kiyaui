<template>
    <v-container class="px-0 margin" fluid>
        <h2 v-if="errorMessage.length > 0">
            {{errorMessage}}
        </h2>
        <v-list>
            <v-subheader>PROFILES</v-subheader>
            <v-list-item-group color="primary" v-model="profile">
                <v-list-item
                    @click="profileSelectionChange(profile)"
                    v-for="profile in profiles" :key="profile">
                    <v-list-item-content>
                        <v-list-item-title v-text="profile"></v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list-item-group>
        </v-list>
    </v-container>
</template>

<script>
    import Wails from '@wailsapp/runtime';
    import { mapState } from 'vuex';

    export default {
        name: "ProfileList.vue",
        data() {
            return {
                profile: 1,
                errorMessage: '',
            }
        },
        computed: mapState({
            profiles: state => state.profiles
        }),
        methods: {
            setErrorMessage: function(message) {
                this.errorMessage = message;
                setTimeout(() => {
                    this.errorMessage = "";
                }, 3000);
            },
            loadProfiles: function () {
                try {
                    this.$store.dispatch('getProfiles');
                } catch (e) {
                    this.setErrorMessage(e.message);
                }
            },
            profileSelectionChange: function (profile) {
                this.$store.dispatch('getProfileKeys', profile);
            }
        },
        mounted() {
            Wails.Events.On("filemodified", () => {
                this.loadProfiles();
            });
            this.loadProfiles();
        }
    }
</script>

<style scoped>
    .margin {
        margin: 16px;
    }
</style>
