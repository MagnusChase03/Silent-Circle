 export default async function useAsymmetricKeys() {
    console.log("Generating RSA key pair...");
  // Generate a new RSA key pair
  const keyPair = await window.crypto.subtle.generateKey(
    {
      name: "RSA-OAEP",
      modulusLength: 2048,
      publicExponent: new Uint8Array([1, 0, 1]),
      hash: "SHA-256",
    },
    true,
    ["encrypt", "decrypt"]
  );

  // Convert the public key and private key to base64 string which is easier to send over the network or store in a database
  const publicKeyBase64 = btoa(
    String.fromCharCode(
      ...new Uint8Array(
        await window.crypto.subtle.exportKey("spki", keyPair.publicKey)
      )
    )
  );
  const privateKeyBase64 = btoa(
    String.fromCharCode(
      ...new Uint8Array(
        await window.crypto.subtle.exportKey("pkcs8", keyPair.privateKey)
      )
    )
  );

  // Export the public key and private key
  return { publicKeyBase64: publicKeyBase64, privateKeyBase64: privateKeyBase64 };
};

