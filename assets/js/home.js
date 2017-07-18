let app = new Vue({
    el: "#main",
    template: `
        <div>
            <h1>Home</h1>
    
            <form @submit.prevent="openRoom">
                <label for="roomname">Create room:</label>
                <input id="roomname" type="text" name="room" v-model="name" @keyUp="checkName" />
                <span :value="error"></span>
            </form>
        </div>
    `,
    data: () => ({
        name: '',
        error: ''
    }),
    methods: {
        checkName() {
            if (!/[a-z0-9_-]/.test(this.name)) {
                this.error = 'Invalid room name';
                return false;
            }
            return true;
        },

        openRoom() {
            if (!this.checkName()) {
                return;
            }
            window.location.href = '/room/' + this.name;
        }
    }
});
