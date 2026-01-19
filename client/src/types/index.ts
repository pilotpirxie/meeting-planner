export interface Calendar {
  id: string;
  title: string;
  description?: string;
  location?: string;
  acceptResponsesUntil?: string;
  createdAt: string;
  updatedAt: string;
}

export interface TimeSlot {
  id: string;
  startDate: string;
  endDate: string;
  createdAt: string;
  updatedAt: string;
}
