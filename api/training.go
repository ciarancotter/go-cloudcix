package api

type CLASS struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		FinishDate string `json:"finish_date"`
		ID         int    `json:"id"`
		StartDate  string `json:"start_date"`
		Syllabus   struct {
			Description string `json:"description"`
			ID          int    `json:"id"`
			MemberID    int    `json:"member_id"`
			Name        string `json:"name"`
			URI         string `json:"uri"`
		} `json:"syllabus"`
		Trainer string `json:"trainer"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type CLASS_SPECIFIC struct {
	Content struct {
		FinishDate string `json:"finish_date"`
		ID         int    `json:"id"`
		StartDate  string `json:"start_date"`
		Syllabus   struct {
			Description string `json:"description"`
			ID          int    `json:"id"`
			MemberID    int    `json:"member_id"`
			Name        string `json:"name"`
			URI         string `json:"uri"`
		} `json:"syllabus"`
		Trainer string `json:"trainer"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type STUDENT struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Cls struct {
			FinishDate string `json:"finish_date"`
			ID         int    `json:"id"`
			StartDate  string `json:"start_date"`
			Syllabus   struct {
				Description string `json:"description"`
				ID          int    `json:"id"`
				MemberID    int    `json:"member_id"`
				Name        string `json:"name"`
				URI         string `json:"uri"`
			} `json:"syllabus"`
			Trainer string `json:"trainer"`
			URI     string `json:"uri"`
		} `json:"cls"`
		ID     int    `json:"id"`
		Notes  string `json:"notes"`
		URI    string `json:"uri"`
		UserID int    `json:"user_id"`
	} `json:"content"`
}

type STUDENT_SPECIFIC struct {
	Content struct {
		Cls struct {
			FinishDate string `json:"finish_date"`
			ID         int    `json:"id"`
			StartDate  string `json:"start_date"`
			Syllabus   struct {
				Description string `json:"description"`
				ID          int    `json:"id"`
				MemberID    int    `json:"member_id"`
				Name        string `json:"name"`
				URI         string `json:"uri"`
			} `json:"syllabus"`
			Trainer string `json:"trainer"`
			URI     string `json:"uri"`
		} `json:"cls"`
		ID     int    `json:"id"`
		Notes  string `json:"notes"`
		URI    string `json:"uri"`
		UserID int    `json:"user_id"`
	} `json:"content"`
}

type SYLLABUS struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Description string `json:"description"`
		ID          int    `json:"id"`
		MemberID    int    `json:"member_id"`
		Name        string `json:"name"`
		URI         string `json:"uri"`
	} `json:"content"`
}

type SYLLABUS_SPECIFIC struct {
	Content struct {
		Description string `json:"description"`
		ID          int    `json:"id"`
		MemberID    int    `json:"member_id"`
		Name        string `json:"name"`
		URI         string `json:"uri"`
	} `json:"content"`
}
