<template>
    <div id="create-user-wrapper">
        <h1>Create Account</h1>
        <form @submit.prevent="createUser">
            <div>
                <label for="username">Username1:</label>
                <input type="text" id="username" v-model="username" required>
            </div>
            <div>
                <label for="password">Password2:</label>
                <input type="password" id="password" v-model="password" required>
            </div>
            <button type="submit">Create Account123</button>
        </form>
    </div>
</template>
<script>
    import { ref } from 'vue';
    import { useRouter } from 'vue-router';
    export default {
    setup() {
        const username = ref('');
        const password = ref('');
        const publicKey = ref(null);
        const router = useRouter();

        async function generateKeyPair() {
            const keyPair = await window.crypto.subtle.generateKey(
                {
                    name: "RSA-OAEP",
                    modulusLength: 2048,
                    publicExponent: new Uint8Array([1, 0, 1]),
                    hash: "SHA-256"
                },
                true,
                ["encrypt", "decrypt"]
            );

            const exportedPublicKey = await window.crypto.subtle.exportKey(
                "spki",
                keyPair.publicKey
            );

            publicKey.value = exportedPublicKey;
        }

        async function createUser() {
            await generateKeyPair();

            const publicKeyBase64 = btoa(String.fromCharCode(...new Uint8Array(publicKey.value)));

            const response = await fetch('/api/create-user', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    username: username.value,
                    password: password.value,
                    publicKey: publicKeyBase64
                })
            });

            if (response.ok) {
                router.push('/login');
            } else {
                console.error('Failed to create user');
            }
        }

        return {
            username,
            password,
            createUser
        };
    }
};
</script>
<style scoped></style>