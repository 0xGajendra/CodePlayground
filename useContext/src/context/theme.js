import {createContext, useContext} from "react";


// before creating context we can predefine the context while creating the context as shown in the below example
export const ThemeContext = createContext({
    themeMode : "light",
    darkTheme : () => {},
    lightTheme : () => {},
}) 

//Whenever some one calles the ThemeContext variable they can access both themeMode and darkTheme/lightTheme variable

// ThemeProvider is a provider component.
// Any component wrapped inside <ThemeProvider> will have access to ThemeContext.
export const ThemeProvider = ThemeContext.Provider;

//creating a custom hook for ThemeContext using useContext
// It calls useContext(ThemeContext), which retrieves the current context value.
export default function useTheme() {
    return useContext(ThemeContext)
}