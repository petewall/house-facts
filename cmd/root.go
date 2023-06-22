package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/petewall/house-facts/internal"
)

var rootCmd = &cobra.Command{
	Use:   "house-facts",
	Short: "house-facts",
	Long:  `house-facts has facts about a house`,
	RunE: func(cmd *cobra.Command, args []string) error {
		facts, err := internal.LoadFacts(viper.GetString("facts-file"))
		if err != nil {
			return err
		}

		r := chi.NewRouter()
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			encoded, err := json.Marshal(facts)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			_, err = w.Write(encoded)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
		})
		r.Get("/metrics", promhttp.Handler().ServeHTTP)

		return http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port")), r)
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("facts-file", "f", "./facts.json", "Path to facts file")
	rootCmd.PersistentFlags().IntP("port", "p", 3000, "Port number to listen on")
	viper.BindPFlag("facts-file", rootCmd.PersistentFlags().Lookup("facts-file"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
