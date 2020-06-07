import kiya from './kiya'

import {
    GET_PROFILES,
    API_FAILURE,
    GET_PROFILE_KEYS,
    API_DATA_PENDING,
    API_DATA_SUCCESS,
    API_DATA_FAILURE,
    SET_ACTIVE_PROFILE
} from './mutation-types'

const actions = {
    getProfiles({ commit }) {
        return kiya.getProfiles()
            .then((response) => commit(GET_PROFILES, response))
            .catch((error) => commit(API_FAILURE, error));
    },
    getProfileKeys({ commit }, payload) {
        commit(SET_ACTIVE_PROFILE, payload);
        commit(API_DATA_PENDING);
        return kiya.getProfileKeys(payload)
            .then((response) => {
                commit(GET_PROFILE_KEYS, response);
                commit(API_DATA_SUCCESS);
            })
            .catch((error) => commit(API_DATA_FAILURE, error));
    }
};

export default actions;
