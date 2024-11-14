import useImportSymmetric from "./useImportSymmetric";

export default async function useEncryptSymMsg(symmetricKeyBase64, message) {
  // import the key from base64 string
  const importedKey = await useImportSymmetric(symmetricKeyBase64);

  // encode the message
  const encodedMessage = new TextEncoder().encode(message);

  // generate an initialization vector with initial value = 0 for simplicity
  const iv = new Uint8Array(16);

  // encrypt the message
  const encrypted = await window.crypto.subtle.encrypt(
    {
      name: "AES-CBC", // this is the encryption algorithm we will use
      iv: iv, // this is the initialization vector we will use
    },
    importedKey, // this is the imported key we will use
    encodedMessage // this is the message we will encrypt
  );

  // encode the encrypted message to base64 which can be stored or transmitted
  return btoa(String.fromCharCode(...new Uint8Array(encrypted)));
}
