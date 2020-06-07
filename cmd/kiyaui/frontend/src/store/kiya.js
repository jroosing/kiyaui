export default {
    async getProfiles() {
        try {
            return await window.backend.Kiya.ProfileNames();
        } catch (e) {
            return e;
        }
    },
    async getProfileKeys(profileName) {
        try {
            return await window.backend.Kiya.ListDetails(profileName)
        } catch (e) {
            return e;
        }
    }
}
