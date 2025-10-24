import { useState } from "react";
import { AppSidebar } from "@/components/AppSidebar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card } from "@/components/ui/card";
import { Send, Loader2 } from "lucide-react";
import { ScrollArea } from "@/components/ui/scroll-area";
import aiAvatar from "@/assets/ai-avatar.jpg";
import { authFetch } from "@/lib/utils";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

interface Message {
  id: number;
  text: string | AIResponse;
  sender: "user" | "ai";
  timestamp: Date;
}

interface AIResponse {
  summary: string;
  recommendations: string;
  disclaimer: string;
}

const Dashboard = () => {
  const [messages, setMessages] = useState<Message[]>([
    {
      id: 1,
      text: "Hello! I'm Health Harbor AI, your personal health advisor. How can I assist you with your health concerns today?",
      sender: "ai",
      timestamp: new Date(),
    },
  ]);
  const [inputMessage, setInputMessage] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const formatAISection = (text: string) => {
    // Replace '*' bullets with '-' for consistency
    const cleaned = text.replace(/\*+/g, "-");

    // Split into lines by either '\n' or '- ' bullet points
    const lines = cleaned
      .split(/\n|-\s+/)
      .map((line) => line.trim())
      .filter(Boolean);

    return lines;
  };

  const parseAIResponse = (text: string): AIResponse => {
    // Use optional hyphen `-?` and trim whitespace around the header
    // This makes the parser more flexible
    const summaryMatch = text.match(
      /-?\s*Summary:([\s\S]*?)(?=-?\s*Recommendations:|$)/
    );
    const recommendationsMatch = text.match(
      /-?\s*Recommendations:([\s\S]*?)(?=This information is|$)/
    );
    const disclaimerMatch = text.match(/This information is[\s\S]*/);

    // Clean up the recommendations by removing the leading '*' from each line
    // react-markdown will add the bullet points back correctly.
    const recommendationsText = recommendationsMatch
      ? recommendationsMatch[1].replace(/^\s*\*\s*/gm, "").trim()
      : "";

    return {
      summary: summaryMatch ? summaryMatch[1].trim() : "No summary provided.",
      recommendations: recommendationsText,
      disclaimer: disclaimerMatch ? disclaimerMatch[0].trim() : "",
    };
  };

  const handleSendMessage = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!inputMessage.trim() || isLoading) return;

    const userMessage: Message = {
      id: messages.length + 1,
      text: inputMessage,
      sender: "user",
      timestamp: new Date(),
    };

    setMessages((prev) => [...prev, userMessage]);
    setInputMessage("");
    setIsLoading(true);

    try {
      // const response = await authFetch("http://localhost:8001/api/chat", {
      //   method: "POST",
      //   headers: { "Content-Type": "application/json" },
      //   body: JSON.stringify({ message: userMessage.text }),
      // });

      // if (!response.ok) {
      //   throw new Error(`Server error: ${response.status}`);
      // }

      // const data = await response.json();

      const data = `
      **Yummy Foods Analysis**

- Summary: The provided foods vary significantly in nutritional content and calorie density. Sweet and sour sauces are high in sugars, while the SlimFast smoothie is high in protein. Yuca cassava chips and turkey salami are lower in calories.

- Recommendations:
    *   Be mindful of portion sizes with sweet and sour sauces due to their high sugar content.
    *   Consider the SlimFast smoothie as a protein source if needed.
    *   Yuca cassava chips and turkey salami can be included in moderation, but prioritize whole, unprocessed foods.

This information is for general knowledge only and does not constitute medical advice. Consult with a qualified healthcare professional or registered dietitian for personalized advice."
`;
      console.log("## Data" + data);

      // Parse the text if it's a JSON string
      // let aiResponseText = data.aiMessage.text;
      let aiResponseText = data;

      // Try to parse if it's a JSON string
      try {
        const parsed = JSON.parse(aiResponseText);
        aiResponseText = parsed.answer || aiResponseText;
      } catch {
        // If not JSON, use as is
      }

      console.log(aiResponseText);

      const aiMessage: Message = {
        id: messages.length + 2,
        text: parseAIResponse(aiResponseText),
        sender: "ai",
        timestamp: new Date(),
      };

      console.log(aiMessage.text);
      if (!aiMessage.text) aiMessage.text = "I can't answer that";

      setMessages((prev) => [...prev, aiMessage]);
    } catch (err: any) {
      const aiMessage: Message = {
        id: messages.length + 2,
        text: "Sorry, something went wrong. Please try again later.",
        sender: "ai",
        timestamp: new Date(),
      };
      setMessages((prev) => [...prev, aiMessage]);
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="flex min-h-screen w-full bg-background">
      <AppSidebar />

      <main className="flex-1 flex flex-col">
        {/* Header */}
        <header className="border-b border-border bg-card">
          <div className="flex items-center justify-between p-6">
            <div>
              <h1 className="text-3xl font-bold text-foreground">
                Health Dashboard
              </h1>
              <p className="text-muted-foreground mt-1">
                Ask me anything about your health
              </p>
            </div>
            <img
              src={aiAvatar}
              alt="AI Assistant Avatar"
              className="h-16 w-16 rounded-full object-cover ring-2 ring-primary"
            />
          </div>
        </header>

        {/* Chat Area */}
        <ScrollArea className="flex-1 p-6">
          <div className="max-w-4xl mx-auto space-y-4">
            {messages.map((message) => (
              <div
                key={message.id}
                className={`flex ${
                  message.sender === "user" ? "justify-end" : "justify-start"
                }`}
              >
                <Card
                  className={`max-w-[80%] p-4 ${
                    message.sender === "user"
                      ? "bg-primary text-primary-foreground"
                      : "bg-card"
                  }`}
                >
                  {typeof message.text === "string" ? (
                    <p className="text-sm leading-relaxed whitespace-pre-line">
                      {message.text}
                    </p>
                  ) : (
                    <div className="text-sm leading-relaxed space-y-4">
                      {/* Summary */}
                      {message.text.summary && (
                        <div>
                          <p className="font-semibold text-base mb-2">
                            Summary
                          </p>
                          <ul className="list-disc ml-5 space-y-1">
                            {formatAISection(message.text.summary).map(
                              (line, idx) => (
                                <li key={idx}>{line}</li>
                              )
                            )}
                          </ul>
                        </div>
                      )}

                      {/* Recommendations */}
                      {message.text.recommendations.length > 0 && (
                        <div>
                          <p className="font-semibold text-base mb-2">
                            Recommendations
                          </p>
                          <ul className="list-disc ml-5 space-y-1">
                            {message.text.disclaimer}
                          </ul>
                        </div>
                      )}

                      {/* Disclaimer */}
                      {message.text.disclaimer && (
                        <p className="text-sm text-muted-foreground mt-4 italic">
                          {message.text.disclaimer}
                        </p>
                      )}
                    </div>
                  )}

                  <p
                    className={`text-xs mt-2 ${
                      message.sender === "user"
                        ? "text-primary-foreground/70"
                        : "text-muted-foreground"
                    }`}
                  >
                    {message.timestamp.toLocaleTimeString([], {
                      hour: "2-digit",
                      minute: "2-digit",
                    })}
                  </p>
                </Card>
              </div>
            ))}
            {isLoading && (
              <div className="flex justify-start">
                <Card className="max-w-[80%] p-4 bg-card">
                  <div className="flex items-center gap-2 text-muted-foreground">
                    <Loader2 className="h-4 w-4 animate-spin" />
                    <p className="text-sm">Health Harbor AI is thinking...</p>
                  </div>
                </Card>
              </div>
            )}
          </div>
        </ScrollArea>

        {/* Input Area */}
        <div className="border-t border-border bg-card p-6">
          <form onSubmit={handleSendMessage} className="max-w-4xl mx-auto">
            <div className="flex gap-3">
              <Input
                value={inputMessage}
                onChange={(e) => setInputMessage(e.target.value)}
                placeholder="Describe your symptoms or ask a health question..."
                className="flex-1 focus-fade"
                disabled={isLoading}
              />
              <Button
                type="submit"
                disabled={isLoading || !inputMessage.trim()}
                variant="accent"
                size="lg"
              >
                <Send className="h-4 w-4" />
              </Button>
            </div>
            <p className="text-xs text-muted-foreground mt-2 text-center">
              This AI advisor provides general health information only. Always
              consult a healthcare professional for medical advice.
            </p>
          </form>
        </div>
      </main>
    </div>
  );
};

export default Dashboard;
