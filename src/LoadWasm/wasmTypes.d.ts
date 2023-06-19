declare global {
    export interface Window {
        Go: any;
        watch: (googleId: string) => any;
        onDataRecieved: (data: any) => any;
    }
}

export { };
