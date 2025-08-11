import fs from "fs";

export default defineEventHandler((_) => {
    const path = "./public/img";
    try {
        const files = shuffle(fs.readdirSync(path));
        return files.map((f: string): string => `/img/${f}`);
    } catch (err) {
        console.error("error reading directory: ", err);
    }
});

// yoink: https://stackoverflow.com/a/2450976
function shuffle(array: string[]): string[] {
    let i: number = array.length;

    while (i != 0) {
        const rand: number = Math.floor(Math.random() * i);
        i--;

        [array[i], array[rand]] = [array[rand], array[i]];
    }

    return array;
}
