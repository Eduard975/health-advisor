import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { toast } from "@/hooks/use-toast";
import defaultImage from "@/assets/default.jpg";

const AuthPage = () => {
  const navigate = useNavigate();
  const [mode, setMode] = useState<"login" | "register">("login");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [fullName, setFullName] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    const endpoint = mode === "login" ? "login" : "register";

    try {
      const response = await fetch(`http://localhost:8001/api/auth/${endpoint}`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(mode === "login" ? { email, password } : { email, password, fullName }),
      });

      const data = await response.json();

      if (!response.ok) {
        toast({
          title: mode === "login" ? "Login failed" : "Registration failed",
          description: data.error || "Something went wrong",
        });
        setIsLoading(false);
        return;
      }

      // Store JWT token
      localStorage.setItem("token", data.token);

      toast({
        title: mode === "login" ? "Welcome back!" : "Registration successful",
        description: `Logged in as ${data.user.email}`,
      });

      navigate("/dashboard");
    } catch (err) {
      console.error(err);
      toast({
        title: "Error",
        description: "Something went wrong. Please try again.",
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex">
      {/* Left side image */}
      <div className="hidden lg:flex lg:w-1/2 relative overflow-hidden">
        <img
          src={defaultImage}
          alt="Medical consultation"
          className="object-cover w-full h-full"
        />
        <div className="absolute inset-0 bg-primary/20 backdrop-blur-[2px]"></div>
        <div className="absolute inset-0 flex flex-col justify-center items-center text-white p-12">
          <h1 className="text-5xl font-bold mb-4">Health Harbor</h1>
          <p className="text-xl text-center max-w-md">
            Your intelligent health companion for personalized medical guidance and support
          </p>
        </div>
      </div>

      {/* Right side - Form */}
      <div className="flex-1 flex items-center justify-center p-8 bg-background">
        <Card className="w-full max-w-md">
          <CardHeader className="space-y-1">
            <CardTitle className="text-3xl font-bold">
              {mode === "login" ? "Welcome back" : "Create an account"}
            </CardTitle>
            <CardDescription>
              {mode === "login"
                ? "Enter your credentials to access your dashboard"
                : "Fill in the form to register"}
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-4">
              {mode === "register" && (
                <div className="space-y-2">
                  <Label htmlFor="fullName">Full Name</Label>
                  <Input
                    id="fullName"
                    type="text"
                    placeholder="John Doe"
                    value={fullName}
                    onChange={(e) => setFullName(e.target.value)}
                    required
                    className="focus-fade"
                  />
                </div>
              )}

              <div className="space-y-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  placeholder="your.email@example.com"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                  className="focus-fade"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <Input
                  id="password"
                  type="password"
                  placeholder="••••••••"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                  className="focus-fade"
                />
              </div>

              <Button type="submit" className="w-full" size="lg" disabled={isLoading}>
                {isLoading ? (mode === "login" ? "Signing in..." : "Registering...") : mode === "login" ? "Sign in" : "Register"}
              </Button>

              <div className="text-center text-sm text-muted-foreground mt-2">
                {mode === "login" ? (
                  <>
                    Don't have an account?{" "}
                    <button
                      type="button"
                      className="text-accent hover:underline"
                      onClick={() => setMode("register")}
                    >
                      Sign up
                    </button>
                  </>
                ) : (
                  <>
                    Already have an account?{" "}
                    <button
                      type="button"
                      className="text-accent hover:underline"
                      onClick={() => setMode("login")}
                    >
                      Sign in
                    </button>
                  </>
                )}
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default AuthPage;
