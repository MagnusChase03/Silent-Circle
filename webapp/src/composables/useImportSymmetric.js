export default async function useImportSymmetric(base64Key) {
  // decode the base64 key
  const binaryKey = Uint8Array.from(atob(base64Key), (c) => c.charCodeAt(0));
  // import the key
  return await window.crypto.subtle.importKey(
    "raw",
    binaryKey.buffer,
    {
      name: "AES-CBC",
    },
    true,
    ["encrypt", "decrypt"]
  );
}
