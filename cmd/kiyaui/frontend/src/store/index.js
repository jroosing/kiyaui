import Vue from "vue";
import Vuex from "vuex";
import actions from './actions';
import {
    GET_PROFILES,
    CHANGE_SELECTED_PROFILE,
    CHANGE_LOADING,
    API_FAILURE,
    GET_PROFILE_KEYS,
    API_DATA_PENDING,
    API_DATA_FAILURE,
    API_DATA_SUCCESS,
    SET_ACTIVE_PROFILE
} from "./mutation-types";

Vue.use(Vuex);

const state = {
    profiles: [],
    profileKeys: [],
    selectedProfile: '',
    errorMessage: '',
    loading: false,
};

const mutations = {
    [GET_PROFILES] (state, payload) {
        this.state.profiles = payload;
    },
    [CHANGE_SELECTED_PROFILE] (state, payload) {
        this.state.selectedProfile = payload.profile;
    },
    [CHANGE_LOADING] (state, payload) {
        this.state.loading = payload.loading;
    },
    [API_FAILURE] (state, payload) {
        this.state.errorMessage = payload.message;
    },
    [GET_PROFILE_KEYS] (state, payload) {
        this.state.profileKeys = payload;
    },
    [API_DATA_PENDING] () {
        this.state.loading = true
    },
    [API_DATA_FAILURE] () {
        this.state.loading = false;
    },
    [API_DATA_SUCCESS] () {
        this.state.loading = false;
    },
    [SET_ACTIVE_PROFILE] (state, payload) {
        this.state.selectedProfile = payload;
    },
    changeErrorMessage(state, payload) {
        this.state.errorMessage = payload.errorMessage;
    },
};

const getters = {
    loading: state => state.loading,
    profiles: state => state.profiles,
    selectedProfile: state => state.selectedProfile,
    errorMessage: state => state.errorMessage,
};

export default new Vuex.Store({
   state,
   getters,
   actions,
   mutations
});
