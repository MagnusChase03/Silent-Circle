<template>
    <div>
        <h2>Encrypt and Decrypt Message</h2>

        <button @click="generateKeyPair">Generate Key Pair</button>
        <div>
            <h3>Public Key:</h3>
            <textarea v-model="publicKeyBase64"></textarea>
        </div>
        <div>
            <h3>Private Key:</h3>
            <textarea v-model="privateKeyBase64"></textarea>
        </div>

        <input type="text" v-model="message" placeholder="Enter message to encrypt"><br />
        <button @click="encryptMessage">Encrypt</button>
        <p>Encrypted Message: <input type="text" v-model="encryptedMessage"></p>
        <button @click="decryptMessage">Decrypt</button>
        <p>Decrypted Message: <input type="text" v-model="decryptedMessage"></p>
    </div>
</template>

<script>
import { ref } from 'vue';
import useAsymmetricKeys from '../composables/useAsymmetricKeys.js';
import useEncryptAsymMsg from '../composables/useEncryptAsymMsg.js';
import useDecryptAsymMsg from '../composables/useDecryptAsymMsg.js';

export default {
    name: 'AsymmetricDemo',
    setup(){
        const publicKeyBase64 = ref('');
        const privateKeyBase64 = ref('');

        const message = ref('');
        const encryptedMessage = ref('');
        const decryptedMessage = ref('');

        async function generateKeyPair() {
            // Generate a new key pair
            const keyPair = await useAsymmetricKeys();
            publicKeyBase64.value = keyPair.publicKeyBase64;
            privateKeyBase64.value = keyPair.privateKeyBase64;
        }

        async function encryptMessage() {
            encryptedMessage.value = await useEncryptAsymMsg(publicKeyBase64.value, message.value);
        }

        async function decryptMessage() {
            decryptedMessage.value = await useDecryptAsymMsg(privateKeyBase64.value, encryptedMessage.value);
        }

        return {
            message,
            encryptedMessage,
            decryptedMessage,
            publicKeyBase64,
            privateKeyBase64,
            generateKeyPair,
            encryptMessage,
            decryptMessage
        };
    }
}
</script>

<style scoped>
textarea,input[type="text"] {
    width: 95%;
    height: 100px;
}
</style>