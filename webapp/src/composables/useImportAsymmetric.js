export default async function useImportAsymmetric(base64Key, type) {
    const binaryKey = Uint8Array.from(atob(base64Key), c => c.charCodeAt(0));
    const format = type === 'public' ? 'spki' : 'pkcs8';
    return await window.crypto.subtle.importKey(
        format,
        binaryKey.buffer,
        {
            name: "RSA-OAEP",
            hash: "SHA-256"
        },
        true,
        type === 'public' ? ["encrypt"] : ["decrypt"]
    ); 
}