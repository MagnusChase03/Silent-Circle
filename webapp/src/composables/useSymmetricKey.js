import { ref } from "vue";

export default async function useSymmetricKey() {
    const symmetricKey = ref(null);
    symmetricKey.value = await window.crypto.subtle.generateKey(
      {
        name: "AES-CBC", // this is the encryption algorithm we will use
        length: 256, // this is the length of the key we will use
      },
      true, // this is whether the key is extractable
      ["encrypt", "decrypt"] // this is whether we can use the key for encryption/decryption
    );

    // export the key to base64 string
    const exportedKey = await window.crypto.subtle.exportKey(
      "raw",
      symmetricKey.value
    );

    // encode the key to base64 which can be stored or transmitted
    return btoa(String.fromCharCode(...new Uint8Array(exportedKey))); 
}