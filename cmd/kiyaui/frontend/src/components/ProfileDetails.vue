<template>
    <v-card class="margin-sides flex-sm-fill">
        <v-spacer></v-spacer>
        <v-card-title>
            {{ selectedProfile }}
            <v-spacer></v-spacer>
            <v-text-field v-model="search" append-icon="mdi-magnify" label="search" single-line hide-details>
            </v-text-field>
        </v-card-title>
        <v-data-table class="flex-sm-fill" :headers="headers" :items="profileKeys" :search="search" :loading="loading" loading-text="Loading kiya data... Please wait">
            <template v-slot:item.actions="{item}">
                <v-tooltip bottom>
                    <template v-slot:activator="{ on }">
                        <v-icon v-on="on" small class="mr-2" @click="copyItem(item)">mdi-clipboard-multiple</v-icon>
                    </template>
                    <span>Copy</span>
                </v-tooltip>
                <v-tooltip bottom>
                    <template v-slot:activator="{ on }">
                        <v-icon v-on="on" small class="mr-2" @click="pasteItem(item)">mdi-clipboard</v-icon>
                    </template>
                    <span>Paste</span>
                </v-tooltip>
                <v-tooltip bottom>
                    <template v-slot:activator="{ on }">
                        <v-icon v-on="on" small class="mr-2" @click="deleteItem(item)">mdi-delete</v-icon>
                    </template>
                    <span>Delete</span>
                </v-tooltip>
            </template>
        </v-data-table>
    </v-card>
</template>

<script>
    import {mapState} from "vuex";

    export default {
        name: "ProfileDetails.vue",
        data: () => ({
            search: '',
            dialog: false,
            headers: [
                {
                    text: 'Name',
                    align: 'start',
                    sortable: true,
                    value: 'Name'
                },
                { text: "Actions", value: "actions", sortable: false }
            ],
        }),
        computed: mapState({
            selectedProfile: state => state.selectedProfile,
            profileKeys: state => state.profileKeys,
            loading: state => state.loading
        }),
        methods: {
            deleteItem(item) {
                const index = this.kiyaKeys.indexOf(item);
                const selectedItem = this.kiyaKeys[index];
                let confirmMsg = "Are you sure you want to delete " + selectedItem.Name + "?";
                confirm(confirmMsg);
            },
        },
    }
</script>

<style scoped>
    .margin-sides {
        margin-left: 16px;
        margin-right: 16px;
    }
</style>
