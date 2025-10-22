import { AppSidebar } from "@/components/AppSidebar";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { Badge } from "@/components/ui/badge";
import { Activity as ActivityIcon, Heart, Footprints, Droplets, Moon, Plus } from "lucide-react";
import { Button } from "@/components/ui/button";

// Mock activity data
const todayStats = {
  steps: 8432,
  stepsGoal: 10000,
  heartRate: 72,
  heartRateStatus: "normal",
  water: 6,
  waterGoal: 8,
  sleep: 7.5,
  sleepGoal: 8,
};

const recentActivities = [
  {
    id: 1,
    type: "Exercise",
    title: "Morning Run",
    duration: "30 min",
    calories: 245,
    time: "7:30 AM",
    intensity: "moderate",
  },
  {
    id: 2,
    type: "Meditation",
    title: "Mindfulness Session",
    duration: "15 min",
    calories: 0,
    time: "12:00 PM",
    intensity: "low",
  },
  {
    id: 3,
    type: "Exercise",
    title: "Yoga Practice",
    duration: "45 min",
    calories: 180,
    time: "6:00 PM",
    intensity: "low",
  },
];

const Activity = () => {
  const stepsProgress = (todayStats.steps / todayStats.stepsGoal) * 100;
  const waterProgress = (todayStats.water / todayStats.waterGoal) * 100;
  const sleepProgress = (todayStats.sleep / todayStats.sleepGoal) * 100;

  return (
    <div className="flex min-h-screen w-full bg-background">
      <AppSidebar />
      
      <main className="flex-1 overflow-auto">
        {/* Header */}
        <header className="border-b border-border bg-card">
          <div className="p-6">
            <div className="flex items-center gap-4">
              <div className="h-12 w-12 rounded-full bg-primary/10 flex items-center justify-center">
                <ActivityIcon className="h-6 w-6 text-primary" />
              </div>
              <div>
                <h1 className="text-3xl font-bold text-foreground">Daily Activity</h1>
                <p className="text-muted-foreground mt-1">
                  Track your health and wellness metrics
                </p>
              </div>
            </div>
          </div>
        </header>

        <div className="p-6 max-w-7xl mx-auto space-y-6">
          {/* Daily Goals */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            {/* Steps */}
            <Card>
              <CardHeader className="pb-3">
                <div className="flex items-center justify-between">
                  <CardDescription>Steps</CardDescription>
                  <Footprints className="h-5 w-5 text-primary" />
                </div>
              </CardHeader>
              <CardContent>
                <div className="space-y-3">
                  <div className="flex items-baseline gap-2">
                    <CardTitle className="text-3xl">{todayStats.steps.toLocaleString()}</CardTitle>
                    <span className="text-sm text-muted-foreground">
                      / {todayStats.stepsGoal.toLocaleString()}
                    </span>
                  </div>
                  <Progress value={stepsProgress} className="h-2" />
                  <p className="text-xs text-muted-foreground">
                    {Math.round(stepsProgress)}% of daily goal
                  </p>
                </div>
              </CardContent>
            </Card>

            {/* Heart Rate */}
            <Card>
              <CardHeader className="pb-3">
                <div className="flex items-center justify-between">
                  <CardDescription>Heart Rate</CardDescription>
                  <Heart className="h-5 w-5 text-accent" />
                </div>
              </CardHeader>
              <CardContent>
                <div className="space-y-3">
                  <div className="flex items-baseline gap-2">
                    <CardTitle className="text-3xl">{todayStats.heartRate}</CardTitle>
                    <span className="text-sm text-muted-foreground">bpm</span>
                  </div>
                  <Badge variant="secondary" className="bg-accent/10 text-accent">
                    {todayStats.heartRateStatus}
                  </Badge>
                  <p className="text-xs text-muted-foreground">
                    Resting heart rate
                  </p>
                </div>
              </CardContent>
            </Card>

            {/* Water Intake */}
            <Card>
              <CardHeader className="pb-3">
                <div className="flex items-center justify-between">
                  <CardDescription>Hydration</CardDescription>
                  <Droplets className="h-5 w-5 text-primary" />
                </div>
              </CardHeader>
              <CardContent>
                <div className="space-y-3">
                  <div className="flex items-baseline gap-2">
                    <CardTitle className="text-3xl">{todayStats.water}</CardTitle>
                    <span className="text-sm text-muted-foreground">
                      / {todayStats.waterGoal} glasses
                    </span>
                  </div>
                  <Progress value={waterProgress} className="h-2" />
                  <p className="text-xs text-muted-foreground">
                    {Math.round(waterProgress)}% of daily goal
                  </p>
                </div>
              </CardContent>
            </Card>

            {/* Sleep */}
            <Card>
              <CardHeader className="pb-3">
                <div className="flex items-center justify-between">
                  <CardDescription>Sleep</CardDescription>
                  <Moon className="h-5 w-5 text-secondary" />
                </div>
              </CardHeader>
              <CardContent>
                <div className="space-y-3">
                  <div className="flex items-baseline gap-2">
                    <CardTitle className="text-3xl">{todayStats.sleep}</CardTitle>
                    <span className="text-sm text-muted-foreground">
                      / {todayStats.sleepGoal} hrs
                    </span>
                  </div>
                  <Progress value={sleepProgress} className="h-2" />
                  <p className="text-xs text-muted-foreground">
                    Last night's sleep
                  </p>
                </div>
              </CardContent>
            </Card>
          </div>

          {/* Recent Activities */}
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle className="text-2xl">Today's Activities</CardTitle>
                  <CardDescription>Your logged activities and exercises</CardDescription>
                </div>
                <Button variant="accent">
                  <Plus className="mr-2 h-4 w-4" />
                  Add Activity
                </Button>
              </div>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {recentActivities.map((activity) => (
                  <div
                    key={activity.id}
                    className="flex items-center justify-between p-4 bg-muted rounded-lg hover:bg-muted/70 transition-colors"
                  >
                    <div className="flex items-center gap-4">
                      <div className="h-12 w-12 rounded-full bg-primary/10 flex items-center justify-center">
                        <ActivityIcon className="h-6 w-6 text-primary" />
                      </div>
                      <div>
                        <div className="flex items-center gap-2 mb-1">
                          <h3 className="font-medium text-foreground">{activity.title}</h3>
                          <Badge
                            variant="secondary"
                            className={
                              activity.intensity === "high"
                                ? "bg-accent/10 text-accent"
                                : activity.intensity === "moderate"
                                ? "bg-primary/10 text-primary"
                                : "bg-secondary/10 text-secondary"
                            }
                          >
                            {activity.intensity}
                          </Badge>
                        </div>
                        <p className="text-sm text-muted-foreground">
                          {activity.duration} • {activity.calories} cal • {activity.time}
                        </p>
                      </div>
                    </div>
                    <Badge>{activity.type}</Badge>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>

          {/* Weekly Summary */}
          <Card>
            <CardHeader>
              <CardTitle className="text-2xl">Weekly Summary</CardTitle>
              <CardDescription>
                Your progress over the last 7 days. Steps are automatically tracked from your logged activities.
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div className="space-y-2">
                  <p className="text-sm text-muted-foreground">Total Steps</p>
                  <p className="text-2xl font-bold text-primary">67,845</p>
                  <p className="text-xs text-muted-foreground">+12% from last week</p>
                </div>
                <div className="space-y-2">
                  <p className="text-sm text-muted-foreground">Active Minutes</p>
                  <p className="text-2xl font-bold text-secondary">245</p>
                  <p className="text-xs text-muted-foreground">+8% from last week</p>
                </div>
                <div className="space-y-2">
                  <p className="text-sm text-muted-foreground">Calories Burned</p>
                  <p className="text-2xl font-bold text-accent">2,840</p>
                  <p className="text-xs text-muted-foreground">-3% from last week</p>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </main>
    </div>
  );
};

export default Activity;
