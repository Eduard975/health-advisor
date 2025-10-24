import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const authFetch = async (url: string, options: RequestInit = {}) => {
  const token = localStorage.getItem("token");
  const headers = {
    ...(options.headers || {}),
    "Authorization": `Bearer ${token}`,
    "Content-Type": "application/json",
  };

  return fetch(url, { ...options, headers });
};
