import useImportSymmetric from "./useImportSymmetric";

export default async function useDecryptSymMsg(
  symmetricKeyBase64,
  encryptedMessage
) {
  // import the key from base64 string
  const importedKey = await useImportSymmetric(symmetricKeyBase64);

  // decode the encrypted message
  const encryptedArray = Uint8Array.from(atob(encryptedMessage), (c) =>
    c.charCodeAt(0)
  );

  // generate an initialization vector with initial value = 0 for simplicity
  const iv = new Uint8Array(16); // must be the same as the one used to encrypt

  // decrypt the message
  const decrypted = await window.crypto.subtle.decrypt(
    {
      name: "AES-CBC", // this is the encryption algorithm we will use
      iv: iv, // this is the initialization vector we will use
    },
    importedKey, // this is the imported key we will use
    encryptedArray // this is the encrypted message we will decrypt
  );

  // decode the decrypted message
  return new TextDecoder().decode(decrypted);
}
