<template>
    <div>
        <h2>Encrypt and Decrypt Message with Symmetric Key</h2>

        <button @click="generateKey">Generate Symmetric Key</button>
        <div>
            <h3>Symmetric Key:</h3>
            <textarea v-model="symmetricKeyBase64"></textarea>
        </div>

        <input type="text" v-model="message" placeholder="Enter message to encrypt"><br />
        <button @click="encryptMessage">Encrypt</button>
        <p>Encrypted Message: <input type="text" v-model="encryptedMessage"></p>
        <button @click="decryptMessage">Decrypt</button>
        <p>Decrypted Message:
            <input type="text" v-model="decryptedMessage">
        </p>
    </div>
</template>

<script>
import { ref } from 'vue';

import useSymmetricKey from '../composables/useSymmetricKey.js';
import useEncryptSymMsg from '../composables/useEncryptSymMsg.js';
import useDecryptSymMsg from '../composables/useDecryptSymMsg.js';

export default {
    name: 'SymmetricDemo',
    setup(){
        const message = ref('');
        const encryptedMessage = ref('');
        const decryptedMessage = ref('');
        const symmetricKeyBase64 = ref('');

        async function generateKey() {
            symmetricKeyBase64.value = await useSymmetricKey();;
        }

        async function encryptMessage() {
            encryptedMessage.value = await useEncryptSymMsg(symmetricKeyBase64.value, message.value);
        }

        async function decryptMessage() {
            decryptedMessage.value = await useDecryptSymMsg(symmetricKeyBase64.value, encryptedMessage.value);;
        }

        return {
            message,
            encryptedMessage,
            decryptedMessage,
            generateKey,
            encryptMessage,
            decryptMessage,
            symmetricKeyBase64
        };
        
    }
};
</script>

<style scoped>
textarea, input[type="text"] {
    width: 95%;
    height: 100px;
    margin-bottom: 10px;
}
</style>