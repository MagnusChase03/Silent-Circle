import useImportAsymmetric from "./useImportAsymmetric";

export default async function useEncryptAsymMsg(publicKeyBase64, message) {
  // Import the public key
  const importedPublicKey = await useImportAsymmetric(
    publicKeyBase64,
    "public"
  );

  // encode the message
  const encodedMessage = new TextEncoder().encode(message);

  // Encrypt the message
  const encrypted = await window.crypto.subtle.encrypt(
    {
      name: "RSA-OAEP",
    },
    importedPublicKey,
    encodedMessage
  );

  // Convert the encrypted message to base64 string
  return btoa(String.fromCharCode(...new Uint8Array(encrypted)));
}

