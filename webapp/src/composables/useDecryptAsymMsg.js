import useImportAsymmetric from "./useImportAsymmetric";

export default async function useDecryptAsymMsg(privateKeyBase64, encryptedMessage) {
  // Import the private key
  const importedPrivateKey = await useImportAsymmetric(privateKeyBase64, "private");

  // convert encrypted message to Uint8Array
  const encryptedArray = Uint8Array.from(atob(encryptedMessage), (c) =>
    c.charCodeAt(0)
  );

  // Decrypt the message
  const decrypted = await window.crypto.subtle.decrypt(
    {
      name: "RSA-OAEP",
    },
    importedPrivateKey,
    encryptedArray
  );

  // decode the decrypted message
  return new TextDecoder().decode(decrypted);
}